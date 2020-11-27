package tag

import (
	"reflect"
	"strings"

	opt "github.com/minitiz/emptyfield/pkg/options"
)

const (
	fieldTag  = "field"
	jsonTag   = "json"
	omitEmpty = "omitempty"
)

type tagOptions string

func fieldPresent(info reflect.StructTag, tagName string) bool {
	tags := info.Get(tagName)
	return tagOptions(tags).Contains(omitEmpty)
}

// OmitEmptyTag ...
func OmitEmptyTag(info reflect.StructTag, opt *opt.Options) bool {
	if opt.JSONOmitEmpty && fieldPresent(info, jsonTag) {
		return true
	}
	return fieldPresent(info, fieldTag)
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var next string
		i := strings.Index(s, ",")
		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}
		if s == optionName {
			return true
		}
		s = next
	}
	return false
}
