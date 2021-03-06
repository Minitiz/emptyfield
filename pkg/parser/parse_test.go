package parser

import (
	"reflect"
	"testing"

	opt "github.com/minitiz/emptyfield/pkg/options"
	"k8s.io/utils/pointer"
)

func Test_addParentPath(t *testing.T) {
	type args struct {
		parentName string
		childrens  []EmptyValues
	}

	type typeof struct {
		test string
	}
	valueof1 := reflect.ValueOf(typeof{"test1"})
	valueof2 := reflect.ValueOf(typeof{"test2"})
	valueof3 := reflect.ValueOf(typeof{"test3"})

	tests := []struct {
		name string
		args args
		want []EmptyValues
	}{
		{
			name: "test 0 - basic",
			args: args{
				parentName: "T.test",
				childrens:  []EmptyValues{{"var", valueof1}},
			},
			want: []EmptyValues{{"T.test.var", valueof1}},
		},
		{
			name: "test 1 - basic multiple",
			args: args{
				parentName: "T.test1",
				childrens:  []EmptyValues{{"var1", valueof1}, {"var2", valueof2}},
			},
			want: []EmptyValues{{"T.test1.var1", valueof1}, {"T.test1.var2", valueof2}},
		},
		{
			name: "test 2 - basic multiple 3",
			args: args{
				parentName: "T.test2",
				childrens:  []EmptyValues{{"var1", valueof1}, {"var2", valueof2}, {"var3", valueof3}},
			},
			want: []EmptyValues{{"T.test2.var1", valueof1}, {"T.test2.var2", valueof2}, {"T.test2.var3", valueof3}},
		},
		{
			name: "test 3 - basic",
			args: args{
				parentName: "",
				childrens:  []EmptyValues{},
			},
			want: []EmptyValues{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addParentPath(tt.args.parentName, tt.args.childrens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s:\ngot:\n\t`%v`\nwant:\n\t`%v`", tt.name, got, tt.want)
			}
		})
	}
}

func TestGetEmptyValues(t *testing.T) {
	type args struct {
		e     reflect.Value
		infos reflect.StructField
		opt   *opt.Options
	}

	type typeof1 struct {
		field1 string
		field2 int64
		field3 []string
		field4 []int64
		field5 struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		}
	}
	type typeof2 struct {
		field1 string `field:"omitempty"`
		field2 int64
		field3 []string
		field4 []int64
		field5 struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		}
	}
	type typeof3 struct {
		field1 string
		field2 int64
		field3 []string
		field4 []int64
		field5 struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		} `field:"omitempty"`
	}
	type typeof4 struct {
		field1 string
		field2 int64
		field3 []string
		field4 []int64
		field5 uint
		field6 bool
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
	var1 := typeof1{
		field1: "test",
		field2: 42,
		field3: []string{"test", "test2"},
		field4: []int64{42, 43, 44},
		field5: struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		}{
			subfield1: "substruct",
			subfield2: 42,
			subfield3: []string{"test", "test2"},
			subfield4: []int64{123, 432, 234},
		},
	}
	valueof1 := reflect.ValueOf(var1)
	var2 := typeof2{
		field1: "",
		field2: 42,
		field3: []string{"test", "test2"},
		field4: []int64{42, 43, 44},
		field5: struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		}{
			subfield1: "substruct",
			subfield2: 42,
			subfield3: []string{"test", "test2"},
			subfield4: []int64{123, 432, 234},
		},
	}
	valueof2 := reflect.ValueOf(var2)
	var3 := typeof3{
		field1: "test",
		field2: 42,
		field3: []string{"test", "test2"},
		field4: []int64{42, 43, 44},
		field5: struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		}{},
	}
	valueof3 := reflect.ValueOf(var3)
	var4 := typeof1{
		field1: "",
		field2: 42,
		field3: []string{"test", "test2"},
		field4: []int64{42, 43, 44},
		field5: struct {
			subfield1 string
			subfield2 int64
			subfield3 []string
			subfield4 []int64
		}{},
	}
	valueof4 := reflect.ValueOf(var4)
	var5 := typeof4{
		field1: "123",
		field2: 123,
		field3: []string{"test"},
		field4: []int64{123, 123},
		field5: 0,
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
				subfield2: 0,
				subfield3: []string{"qwe"},
				subfield4: []int64{123, 123},
			},
		},
		field10: 4.2,
	}
	valueof5 := reflect.ValueOf(var5)

	tests := []struct {
		name    string
		args    args
		wantRet []EmptyValues
	}{
		{
			name: "test 0 - basic",
			args: args{
				e:     valueof1,
				infos: reflect.StructField{Name: "T"},
				opt:   &opt.Options{},
			},
			wantRet: nil,
		},
		{
			name: "test 1 - omitempty basic",
			args: args{
				e:     valueof2,
				infos: reflect.StructField{Name: "T"},
				opt: &opt.Options{
					Panic: false,
					Tags:  []string{"field"},
				},
			},
			wantRet: nil,
		},
		{
			name: "test 2 - omitempty struct basic",
			args: args{
				e:     valueof3,
				infos: reflect.StructField{Name: "T"},
				opt: &opt.Options{
					Panic: false,
					Tags:  []string{"field"},
				},
			},
			wantRet: nil,
		},
		{
			name: "test 3 - basic",
			args: args{
				e:     valueof4,
				infos: reflect.StructField{Name: "T"},
				opt: &opt.Options{
					Panic: false,
					Tags:  []string{"field"},
				},
			},
			wantRet: []EmptyValues{
				{
					Variable: "T.field1",
				},
				{
					Variable: "T.field5.subfield1",
				},
				{
					Variable: "T.field5.subfield2",
				},
				{
					Variable: "T.field5.subfield3",
				},
				{
					Variable: "T.field5.subfield4",
				},
			},
		},
		{
			name: "test 4 -  basic all case",
			args: args{
				e:     valueof5,
				infos: reflect.StructField{Name: "T"},
				opt: &opt.Options{
					Panic: false,
					Tags:  []string{"field"},
				},
			},
			wantRet: []EmptyValues{
				{
					Variable: "T.field5",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRet := GetEmptyValues(tt.args.e, tt.args.infos, tt.args.opt)
			ret := []string{}
			want := []string{}
			for i := range gotRet {
				ret = append(ret, gotRet[i].Variable)
			}
			for i := range tt.wantRet {
				want = append(want, tt.wantRet[i].Variable)
			}
			if !reflect.DeepEqual(ret, want) {
				t.Errorf("%s:\ngot:\n\t`%v`\nwant:\n\t`%v`", tt.name, gotRet, tt.wantRet)
			}

		})
	}
}
