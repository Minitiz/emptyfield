package tag

import (
	"testing"
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
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Contains(tt.args.optionName); got != tt.want {
				t.Errorf("tagOptions.Contains() = %v, want %v", got, tt.want)
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

// func Test_OmitEmpty(t *testing.T) {

// 	tests := []struct {
// 		name string
// 		data interface{}
// 		opt  *opt.Option
// 		want bool
// 	}{

// 		//presence omitempty on field tag
// 		{"test 1", fieldTagTest{"test"}, &opt.Options{}, true},
// 		{"test 2", jsonTagTest{"test"}, &opt.Options{true, false}, true},
// 		{"test 3", jsonTagTest{"test"}, &opt.Options{}, false},
// 		{"test 4", noTagTest{"test"}, &opt.Options{}, false},
// 	}

// 	//presence omitempty on json tag but opt nil
// 	// {
// 	// 	name: "test 2",
// 	// 	data struct{
// 	// 		test string `json:",omitempty`
// 	// 	}{
// 	// 		"test",
// 	// 	},
// 	// 	opt: nil,
// 	// 	want false,
// 	// },
// 	//presence omitempty on json tag and opt say to parse it

// 	// TODO: Add test cases.

// 	// for _, tt := range tests {
// 	// 	t.Run(tt.name, func(t *testing.T) {
// 	// 		if got := OmitEmptyTag(reflect.TypeOf(tt.data).Field(0).Tag,tt.opt) ; got != tt.want {
// 	// 			t.Errorf("tagOptions.Contains() = %v, want %v", got, tt.want)
// 	// 		}
// 	// 	})
// 	// }
// }
