package result

import (
	"bytes"
	"errors"
)

var VoidValue = &Void{}

var nullJsonValue = []byte("null")

type Void struct {
}

func (s *Void) UnmarshalJSON(b []byte) error {
	if 0 != bytes.Compare(nullJsonValue, b) {
		return errors.New("just parse 'null' string")
	}
	return nil
}

func (s *Void) MarshalJSON() ([]byte, error) {
	return nullJsonValue, nil
}
