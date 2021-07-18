package callback

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"

	"github.com/sfpyhub/webhook-test/definitions"
)

// genMAC generates the HMAC signature for a message provided the secret key
// and hashFunc.
func genMAC(message, key []byte, hashFunc func() hash.Hash) []byte {
	mac := hmac.New(hashFunc, key)
	mac.Write(message)
	return mac.Sum(nil)
}

// checkMAC reports whether messageMAC is a valid HMAC tag for message.
func checkMAC(message, messageMAC, key []byte, hashFunc func() hash.Hash) bool {
	expectedMAC := genMAC(message, key, hashFunc)
	return hmac.Equal(messageMAC, expectedMAC)
}

// messageMAC returns the hex-decoded HMAC tag from the signature and its
// corresponding hash function.
func messageMAC(signature string) ([]byte, error) {
	if signature == "" {
		return nil, errors.New("missing signature")
	}

	buf, err := hex.DecodeString(signature)
	if err != nil {
		return nil, fmt.Errorf("error decoding signature %q: %v", signature, err)
	}
	return buf, nil
}

func ValidateSignature(signature string, event *definitions.Event, secretToken []byte) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	messageMAC, err := messageMAC(signature)
	if err != nil {
		return err
	}
	if !checkMAC(payload, messageMAC, secretToken, sha512.New) {
		return errors.New("payload signature check failed")
	}
	return nil
}
