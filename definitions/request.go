package definitions

import "github.com/sfpyhub/go-sfpy/responses"

// Head represents the headers
type Head struct {
	Signature string `json:"-"`
}

type Request struct {
	Head Head `json:"-" form:"-"`

	Data *responses.Event `json:"data,omitempty" form:"-"`
}
