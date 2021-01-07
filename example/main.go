package main

import (
	"log"
	"reflect"

	"github.com/minitiz/emptyfield"
	opt "github.com/minitiz/emptyfield/pkg/options"
	"k8s.io/utils/pointer"
)

type test struct {
	field1 string
	field2 int64
	field3 []string `yaml:",omitempty"`
	field4 []int64  `json:",omitempty"`
	field5 uint
	field6 bool `field:"omitempty"`
	field7 *string
	field8 map[string]interface{}
	field9 []struct {
		subfield1 string
		subfield2 int64
		subfield3 []string
		subfield4 []int64
	} `field:"omitempty"`
	field10 float64
}

func main() {

	// simple usecase
	var variable = test{
		field1: "123",
		field2: 123,
		field3: []string{"test"},
		field4: []int64{123, 123},
		field5: 64,
		field6: true,
		field7: pointer.StringPtr("test"),
		field8: map[string]interface{}{"test": "test"},
		field9: []struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		}{
			{
				subfield1: "qwe",
				subfield2: 51,
				subfield3: []string{"qwe"},
				subfield4: []int64{123, 123},
			},
		},
		field10: 4.2,
	}
	emptyfield1, err := emptyfield.Check(reflect.ValueOf(variable))
	log.Printf("Example 1 - return emptyfield => \n\t`%v`\n && error => \n\t`%v`\n", emptyfield1, err)

	// usecase with emptyfield
	// field2 will be equal to 0 with no tag omitempty on his field
	// check will return an error
	var variable2 = test{
		field1: "123",
		field2: 0,
		field3: []string{"test"},
		field4: []int64{123, 123},
		field5: 64,
		field6: true,
		field7: pointer.StringPtr("test"),
		field8: map[string]interface{}{"test": "test"},
		field9: []struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		}{
			{
				subfield1: "qwe",
				subfield2: 51,
				subfield3: []string{"qwe"},
				subfield4: []int64{123, 123},
			},
		},
		field10: 4.2,
	}
	emptyfield2, err := emptyfield.Check(reflect.ValueOf(variable2))
	log.Printf("Example 2 - return emptyfield => \n\t`%v`\n && error => \n\t`%v`\n", emptyfield2, err)

	// usecase with emptyfield but omitempty tag is present for this field
	// field6 will be equal to false but omitempty is present on his tag
	var variable3 = test{
		field1: "123",
		field2: 42,
		field3: []string{"test"},
		field4: []int64{123, 123},
		field5: 64,
		field6: false,
		field7: pointer.StringPtr("test"),
		field8: map[string]interface{}{"test": "test"},
		field9: []struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		}{
			{
				subfield1: "qwe",
				subfield2: 51,
				subfield3: []string{"qwe"},
				subfield4: []int64{123, 123},
			},
		},
		field10: 4.2,
	}
	emptyfield3, err := emptyfield.Check(reflect.ValueOf(variable3))
	log.Printf("Example 3 - return emptyfield => \n\t`%v`\n && error => \n\t`%v`\n", emptyfield3, err)

	// usecase with json tag with omitempty declared
	// field 4 have json omitempty but no opt is sended to Check and return an error
	var variable4 = test{
		field1: "123",
		field2: 42,
		field3: []string{"test"},
		field4: []int64{},
		field5: 64,
		field6: false,
		field7: pointer.StringPtr("test"),
		field8: map[string]interface{}{"test": "test"},
		field9: []struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		}{
			{
				subfield1: "qwe",
				subfield2: 51,
				subfield3: []string{"qwe"},
				subfield4: []int64{123, 123},
			},
		},
		field10: 4.2,
	}
	emptyfield4, err := emptyfield.Check(reflect.ValueOf(variable4))
	log.Printf("Example 4 - return emptyfield => \n\t`%v`\n && error => \n\t`%v`\n", emptyfield4, err)

	// usecase with json tag with omitempty declared
	// field 4 have json omitempty and opt for enable json tag parsing is sended to Check
	// no error will be returned this time
	// you can execute the same thing with yaml and test it on field 3
	var variable5 = test{
		field1: "123",
		field2: 42,
		field3: []string{"test"},
		field4: []int64{},
		field5: 64,
		field6: false,
		field7: pointer.StringPtr("test"),
		field8: map[string]interface{}{"test": "test"},
		field9: []struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		}{
			{
				subfield1: "qwe",
				subfield2: 42,
				subfield3: []string{"qwe"},
				subfield4: []int64{123, 123},
			},
		},
		field10: 4.2,
	}
	emptyfield5, err := emptyfield.Check(reflect.ValueOf(variable5), opt.OmitEmptyTag("json"))
	log.Printf("Example 5 - return emptyfield => \n\t`%v`\n && error => \n\t`%v`\n", emptyfield5, err)

	// usecase with one field equal to zero on a struct
	// field9 have tag with omitempty for his entire struct and subfield2 is equal to 0
	// no error will be returned because this tag is present on his parent
	var variable6 = test{
		field1: "123",
		field2: 42,
		field3: []string{"test"},
		field4: []int64{},
		field5: 64,
		field6: false,
		field7: pointer.StringPtr("test"),
		field8: map[string]interface{}{"test": "test"},
		field9: []struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		}{
			{
				subfield1: "qwe",
				subfield2: 0,
				subfield3: []string{"qwe"},
				subfield4: []int64{123, 123},
			},
		},
		field10: 4.2,
	}
	emptyfield6, err := emptyfield.Check(reflect.ValueOf(variable6), opt.OmitEmptyTag("json"))
	log.Printf("Example 6 - return emptyfield => \n\t`%v`\n && error => \n\t`%v`\n", emptyfield6, err)

}
