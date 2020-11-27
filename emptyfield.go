package emptyfield

import (
	"errors"
	"reflect"

	"github.com/google/martian/log"
	opt "github.com/minitiz/emptyfield/pkg/options"
	"github.com/minitiz/emptyfield/pkg/parser"
)

// Empty ...
type Empty []string

// ErrorGeneric is send to panic if opt.Panic is enabled
var ErrorGeneric error = errors.New("Empty fields detected")

func formatReturn(EmptyParsed []parser.EmptyValues, opts *opt.Options) (ret Empty) {
	for i := range EmptyParsed {
		ret = append(ret, EmptyParsed[i].Variable)
	}
	if opts.Panic {
		log.Errorf("%v", ErrorGeneric)
		for i := range ret {
			log.Errorf("%s value is missing", ret[i])
		}
		panic(ErrorGeneric)
	}
	return ret
}

// Check your structure fields and detect empty field into your struct
// return structure filled with fields empty of your go struct
// able to skip an omitempty field if field tag is provided on your structure as `field:"omitempty"`
// able to skip an omitempty field with json tag if JSONOmitEmptyEnabled is past on params
// able to panic if PanicEnabled is past on params.
func Check(data reflect.Value, opts ...opt.Option) Empty {
	f := &opt.Options{}
	// Option paremeters values:
	for _, op := range opts {
		op(f)
	}
	// checkData() on doit checker si data a au moins un field rempli
	return formatReturn(parser.GetEmptyValues(reflect.Indirect(data), reflect.StructField{Name: "T"}, f), f)
}
