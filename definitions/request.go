package definitions

// Head represents the headers
type Head struct {
	Signature string `json:"-"`
}

type Request struct {
	Head Head `json:"-" form:"-"`

	Data *Event `json:"data,omitempty" form:"-"`
}
