package emptyfield

import (
	"reflect"

	"github.com/google/martian/log"
	opt "github.com/minitiz/emptyfield/pkg/options"
	"github.com/minitiz/emptyfield/pkg/parser"
)

// Empty is return to function as result
type Empty []string

// ErrorType ...
type ErrorType string

func (e ErrorType) Error() string {
	return string(e)
}

// ErrorGeneric is send to panic if opt.Panic is enabled
const (
	ErrorGeneric ErrorType = "Empty fields detected"
	ErrorCritic  ErrorType = "Given structure is nil"
)

func formatReturn(EmptyParsed []parser.EmptyValues, opts *opt.Options) (ret Empty, err error) {
	for i := range EmptyParsed {
		ret = append(ret, EmptyParsed[i].Variable)
	}
	if len(EmptyParsed) > 0 {
		for i := range ret {
			log.Errorf("%s value is missing", ret[i])
		}
		if opts.Panic {
			panic(ErrorGeneric)
		}
		err = ErrorGeneric
		return
	}
	return nil, err
}

// Check your structure fields and detect empty field into your struct
// return structure filled with fields empty of your go struct
// able to skip an omitempty field if field tag is provided on your structure as `field:"omitempty"`
// able to skip an omitempty field with json tag if JSONOmitEmptyEnabled is past on params
// able to panic if PanicEnabled is past on params.
func Check(data reflect.Value, opts ...opt.Option) (Empty, error) {
	if !data.IsValid() {
		return nil, ErrorCritic
	}

	f := &opt.Options{}
	for _, op := range opts {
		op(f)
	}
	f.ApplyTagEmptyValue()
	return formatReturn(parser.GetEmptyValues(reflect.Indirect(data), reflect.StructField{Name: "T"}, f), f)
}
