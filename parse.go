package emptyfield

import (
	"fmt"
	"reflect"
)

// emptyValues ...
type emptyValues struct {
	variable string
	ref      reflect.Value
}

func addParentPath(parentName string, childrens []emptyValues) []emptyValues {
	for i := range childrens {
		childrens[i].variable = fmt.Sprintf("%s.%s", parentName, childrens[i].variable)
	}
	return childrens
}

func fieldPresent(info reflect.StructTag, tagName string) bool {
	tags := info.Get(tagName)
	return tagOptions(tags).Contains(omitEmpty)
}

func omitEmptyTag(info reflect.StructTag, opt *Options) bool {
	if opt.JSONOmitEmpty {
		if fieldPresent(info, jsonTag) {
			return true
		}
	}
	return fieldPresent(info, fieldTag)
}

// getEmptyValues return all empty field name with his parents's name
func getEmptyValues(e reflect.Value, infos reflect.StructField, opt *Options) (ret []emptyValues) {
	empty := false
	switch e.Type().Kind() {
	case reflect.Ptr:
		empty = e.IsNil()
		if !empty {
			ret = append(ret, getEmptyValues(e.Elem(), infos, opt)...)
		}
	case reflect.Struct:
		for i := e.NumField() - 1; i >= 0; i-- {
			if EmptyChildren := getEmptyValues(e.Field(i), reflect.TypeOf(e.Interface()).Field(i), opt); len(EmptyChildren) > 0 {
				ret = append(ret, addParentPath(infos.Name, EmptyChildren)...)
			}
		}
	case reflect.Array, reflect.Slice:
		for i := e.Len() - 1; i >= 0; i-- {
			if EmptyChildren := getEmptyValues(e.Index(i), reflect.TypeOf(e.Interface()).Field(i), opt); len(EmptyChildren) > 0 {
				ret = append(ret, addParentPath(infos.Name, EmptyChildren)...)
			}
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
	if empty && !omitEmptyTag(infos.Tag, opt) {
		ret = append(ret, emptyValues{infos.Name, e})
	}
	return ret
}
