package emptyfield

import (
	"reflect"
	"testing"

	"github.com/minitiz/emptyfield"
)

type datatest struct {
	testname string
	data     interface{}
	want     *emptyfield.Empty
}

var tests = makeTest()

func makeTest() []datatest {
	return []datatest{
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
		// 	want: &emptyfield.Empty{
		// 		Fields: []string{"test"},
		// 		Err:    emptyfield.Error,
		// 	},
		// },
		// {
		// 	testname: "test 3",
		// 	data: struct {
		// 		test string
		// 	}{
		// 		"",
		// 	},
		// 	want: &emptyfield.Empty{
		// 		Fields: []string{"test"},
		// 		Err:    emptyfield.Error,
		// 	},
		// },
		// {
		// 	testname: "test 4",
		// 	data: struct {
		// 		test string
		// 	}{
		// 		"",
		// 	},
		// 	want: &emptyfield.Empty{
		// 		Fields: []string{"test"},
		// 		Err:    emptyfield.Error,
		// 	},
		// },
		// {
		// 	testname: "test 5",
		// 	data: struct {
		// 		test string
		// 	}{
		// 		"",
		// 	},
		// 	want: &emptyfield.Empty{
		// 		Fields: []string{"test"},
		// 		Err:    emptyfield.Error,
		// 	},
		// },
	}
}

func TestCheck(t *testing.T) {
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			want := tt.want
			got := emptyfield.Check(tt.data)
			if !reflect.DeepEqual(want, got) {
				t.Errorf("%s:\ngot:\n\t`%v`\nwant:\n\t`%v", tt.testname, want, got)
			}
		})
	}
}

func TestPanic(t *testing.T) {

}
