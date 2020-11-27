package emptyfield

import (
	"reflect"
	"testing"
)

var tests = []struct {
	testname string
	data     interface{}
	want     *Empty
}{
	{
		testname: "test 1",
		data: struct {
			test string
		}{
			"allo",
		},
		want: nil,
	},
	// {
	// 	testname: "test 2",
	// 	data: struct {
	// 		test string
	// 	}{
	// 		"",
	// 	},
	// 	want: Empty{
	// 		Fields: []string{"test"},
	// 		Err:    Error,
	// 	},
	// },
	// {
	// 	testname: "test 3",
	// 	data: struct {
	// 		test string
	// 	}{
	// 		"",
	// 	},
	// 	want: Empty{
	// 		Fields: []string{"test"},
	// 		Err:    Error,
	// 	},
	// },
	// {
	// 	testname: "test 4",
	// 	data: struct {
	// 		test string
	// 	}{
	// 		"",
	// 	},
	// 	want: Empty{
	// 		Fields: []string{"test"},
	// 		Err:    Error,
	// 	},
	// },
	// {
	// 	testname: "test 5",
	// 	data: struct {
	// 		test string
	// 	}{
	// 		"",
	// 	},
	// 	want: Empty{
	// 		Fields: []string{"test"},
	// 		Err:    Error,
	// 	},
	// },
}

func TestCheck(t *testing.T) {
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			want := tt.want
			got := Check(tt.data)
			if !reflect.DeepEqual(want, got) {
				t.Errorf("%s:\ngot:\n\t`%v`\nwant:\n\t`%v", tt.testname, want, got)
			}
		})
	}
}

func TestPanic(t *testing.T) {

}
