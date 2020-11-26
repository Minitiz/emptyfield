package emptyfield

import "errors"

// EmptyField check if one field of your structure is empty
type EmptyField struct {
}

// OmitEmpty ...
func (ef *EmptyField) OmitEmpty(activate bool) {

}

// Empty ...
type Empty struct {
	Fields []string
	Err    error
}

// Option ...
type Option func(*Options)

// Options ...
type Options struct {
	OmitEmpty bool
}

// Error is the generic error return if one or more fields is/are empty
var Error error = errors.New("Empty fields detected")

// Check your structure fields and return each field empty
func Check(data interface{}, opts ...Option) *Empty {

	return nil
}

// Panic parse your structure fields, panic if  (using logrus)
func (ef *EmptyField) Panic(data interface{}, opts ...Option) {

}
