package tag

import (
	"reflect"
	"testing"

	opt "github.com/minitiz/emptyfield/pkg/options"
)

type datatest struct {
	testname string
	opt      string
	want     bool
}

// func TestContains(t *testing.T) {
// 	name, opts := parserTag("fields")
// 	fmt.println(name)
// 	fmt.println(opts)
// }

func Test_Contains(t *testing.T) {
	type args struct {
		optionName string
	}
	tests := []struct {
		name string
		o    tagOptions
		args args
		want bool
	}{
		{
			name: "test 0",
			o:    `omitempty`,
			args: args{omitEmpty},
			want: true,
		},
		{
			name: "test 1",
			o:    `hello,omitempty`,
			args: args{omitEmpty},
			want: true,
		},
		{
			name: "test 2",
			o:    `hello,`,
			args: args{omitEmpty},
			want: false,
		},
		{
			name: "test 3",
			o:    `hello,bonjour`,
			args: args{omitEmpty},
			want: false,
		},
		{
			name: "test 4",
			o:    ``,
			args: args{omitEmpty},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Contains(tt.args.optionName); got != tt.want {
				t.Errorf("%s:\ngot:\n\t`%v`\nwant:\n\t`%v`", tt.name, got, tt.want)
			}
		})
	}
}

type fieldTagTest struct {
	data string `field:"omitempty"`
}

type jsonTagTest struct {
	data string `json:",omitempty"`
}

type noTagTest struct {
	data string
}

type yamlTagTest struct {
	data string `yaml:",omitempty"`
}

func Test_OmitEmpty(t *testing.T) {

	tests := []struct {
		name string
		data interface{}
		opt  *opt.Options
		want bool
	}{

		//presence omitempty on field tag
		{"test 0 - test field tag", fieldTagTest{"test"}, &opt.Options{false, []string{"field"}}, true},
		{"test 1 - test json tag", jsonTagTest{"test"}, &opt.Options{false, []string{"json"}}, true},
		{"test 2 - test json tag", jsonTagTest{"test"}, &opt.Options{false, []string{}}, false},
		{"test 3 - test yaml tag", yamlTagTest{"test"}, &opt.Options{false, []string{"yaml"}}, true},
		{"test 4 - test yaml tag", yamlTagTest{"test"}, &opt.Options{false, []string{}}, false},
		{"test 5 - test no tag", noTagTest{"test"}, &opt.Options{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OmitEmptyTag(reflect.TypeOf(tt.data).Field(0).Tag, tt.opt); got != tt.want {
				t.Errorf("%s:\ngot:\n\t`%v`\nwant:\n\t`%v`", tt.name, got, tt.want)
			}
		})
	}
}
