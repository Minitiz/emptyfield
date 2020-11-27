package emptyfield

import (
	"errors"
	"reflect"
)

// EmptyField check if one field of your structure is empty
type EmptyField struct {
}

// OmitEmpty ...
func (ef *EmptyField) OmitEmpty(activate bool) {

}

// Empty ...
type Empty []string

// Option ...
type Option func(*Options)

// Options ...
type Options struct {
	JSONOmitEmpty bool
	Panic         bool
}

// Error is the generic error return if one or more fields is/are empty
var Error error = errors.New("Empty fields detected")

// JSONOmitEmptyEnabled ...
func JSONOmitEmptyEnabled(f *Options) {
	f.JSONOmitEmpty = true
	return
}

// PanicEnabled allow Check to panic if an empty value is detected
func PanicEnabled(f *Options) {
	f.Panic = true
	return
}

// Data ...
type Data reflect.Value

// Check your structure fields and detect empty field into your struct
// return structure filled with fields empty of your go struct
// able to skip an omitempty field if field tag is provided on your structure as `field:"omitempty"`
// able to skip an omitempty field with json tag if JSONOmitEmptyEnabled is past on params
// able to panic if PanicEnabled is past on params.
func Check(data reflect.Value, opts ...Option) (Empty, error) {
	f := &Options{}
	// Option paremeters values:
	for _, op := range opts {
		op(f)
	}
	var ret []string
	// checkData() on doit checker si data a au moins un field rempli
	EmptyToParse := getEmptyValues(reflect.Indirect(data), reflect.StructField{}, f)
	for i := range EmptyToParse {
		ret = append(ret, EmptyToParse[i].variable)
	}
	// if an omitempty tag is set for this value
	return ret, nil
}
