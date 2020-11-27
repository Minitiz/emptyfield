package parser

import (
	"fmt"
	"reflect"

	opt "github.com/minitiz/emptyfield/pkg/options"
	"github.com/minitiz/emptyfield/pkg/tag"
)

// EmptyValues ...
type EmptyValues struct {
	Variable string
	Ref      reflect.Value
}

// // Options ...
// type Options emptyfield.Options

func addParentPath(parentName string, childrens []EmptyValues) []EmptyValues {
	for i := range childrens {
		childrens[i].Variable = fmt.Sprintf("%s.%s", parentName, childrens[i].Variable)
	}
	return childrens
}

// GetEmptyValues return all empty field name with his parents's name
func GetEmptyValues(e reflect.Value, infos reflect.StructField, opt *opt.Options) (ret []EmptyValues) {
	empty := false
	ChildrenToAdd := []EmptyValues{}
	switch e.Type().Kind() {
	case reflect.Ptr:
		empty = e.IsNil()
		if !empty {
			ret = append(ret, GetEmptyValues(e.Elem(), infos, opt)...)
		}
	case reflect.Struct:
		emptyCounter := 0
		for i := e.NumField() - 1; i >= 0; i-- {
			if EmptyChildren := GetEmptyValues(e.Field(i), reflect.TypeOf(e.Interface()).Field(i), opt); len(EmptyChildren) > 0 {
				ChildrenToAdd = append(ChildrenToAdd, addParentPath(infos.Name, EmptyChildren)...)
				emptyCounter++
			}
			empty = emptyCounter == e.NumField()
		}
	case reflect.Array, reflect.Slice:
		for i := e.Len() - 1; i >= 0; i-- {
			emptyCounter := 0
			if EmptyChildren := GetEmptyValues(e.Index(i), reflect.TypeOf(e.Interface()).Field(i), opt); len(EmptyChildren) > 0 {
				ChildrenToAdd = append(ChildrenToAdd, addParentPath(infos.Name, EmptyChildren)...)
				emptyCounter++
			}
			empty = emptyCounter == e.Len()
		}
	case reflect.String:
		empty = e.Len() == 0
	case reflect.Bool:
		empty = !e.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		empty = e.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		empty = e.Uint() == 0
	case reflect.Float32, reflect.Float64:
		empty = e.Float() == 0
	case reflect.Interface, reflect.Map:
		empty = e.IsNil()
	}
	if empty && !tag.OmitEmptyTag(infos.Tag, opt) {
		ret = append(ret, EmptyValues{infos.Name, e})
	} else if len(ChildrenToAdd) > 0 {
		ret = append(ret, ChildrenToAdd...)
	}
	return ret
}