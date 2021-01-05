package emptyfield

import (
	"reflect"
	"testing"

	"k8s.io/utils/pointer"
)

type ret struct {
	data Empty
	err  error
}

var tests = []struct {
	testname string
	data     interface{}
	want     interface{}
}{
	{
		testname: "test 0 - basic",
		data: &struct {
			test string
		}{
			"allo",
		},
		want: ret{
			nil,
			nil,
		},
	},
	{
		testname: "test 1 - send nil struct to Check",
		data:     nil,
		want: ret{
			nil,
			ErrorCritic,
		},
	},
	{
		testname: "test 2- basic",
		data: &struct {
			str    string
			nbr    int64
			tabstr []string
			tabnbr []int64
		}{
			str:    "crumble",
			nbr:    85000,
			tabstr: []string{"test", "crumble", "kai"},
			tabnbr: []int64{5, 123, 546},
		},
		want: ret{
			nil,
			nil,
		},
	},
	{
		testname: "test 3 - field omitempty basic",
		data: &struct {
			str    string
			nbr    int64
			tabstr []string
			tabnbr []int64 `field:"omitempty"`
		}{
			"crumble",
			85000,
			[]string{"test", "crumble", "kai"},
			nil,
		},
		want: ret{
			nil,
			nil,
		},
	},
	{
		testname: "test 4 - basic",
		data: &struct {
			str    string
			nbr    int64
			tabstr []string
			tabnbr []int64
		}{
			"crumble",
			85000,
			[]string{"test", "crumble", "kai"},
			nil,
		},
		want: ret{
			[]string{"T.tabnbr"},
			ErrorGeneric,
		},
	},
	{
		testname: "test 5 - basic ptr",
		data: &struct {
			str    *string
			nbr    *int64
			tabstr *[]string
			tabnbr *[]int64
		}{
			str:    pointer.StringPtr("test"),
			nbr:    pointer.Int64Ptr(42),
			tabstr: &[]string{"test", "crumble"},
			tabnbr: &[]int64{4, 5, 6},
		},
		want: ret{
			nil,
			nil,
		},
	},
	{
		testname: "test 6 - basic ptr",
		data: &struct {
			str    *string
			nbr    *int64
			tabstr *[]string
			tabnbr *[]int64
		}{
			str:    pointer.StringPtr("test"),
			nbr:    pointer.Int64Ptr(42),
			tabstr: &[]string{"test", "crumble"},
		},
		want: ret{
			[]string{"T.tabnbr"},
			ErrorGeneric,
		},
	},
	{
		testname: "test 7 - field omitempty basic ptr",
		data: &struct {
			str    *string
			nbr    *int64
			tabstr *[]string
			tabnbr *[]int64 `field:"omitempty"`
		}{
			str:    pointer.StringPtr("test"),
			nbr:    pointer.Int64Ptr(42),
			tabstr: &[]string{"test", "crumble"},
		},
		want: ret{
			nil,
			nil,
		},
	},
	{
		testname: "test 8 - complex struct",
		data: &struct {
			data struct {
				str    string
				nbr    int64
				tabstr []string
				tabnbr []int64
			}
		}{
			data: struct {
				str    string
				nbr    int64
				tabstr []string
				tabnbr []int64
			}{
				// str:    "crumble",
				nbr:    4,
				tabstr: []string{"crumble", "kai"},
				tabnbr: []int64{1, 2, 3, 4, 5},
			},
		},
		want: ret{
			[]string{"T.data.str"},
			ErrorGeneric,
		},
	},
	{
		testname: "test 9 - complex struct multiple empty",
		data: &struct {
			data struct {
				str    string
				nbr    int64
				tabstr []string
				tabnbr []int64
			}
		}{
			data: struct {
				str    string
				nbr    int64
				tabstr []string
				tabnbr []int64
			}{
				nbr:    4,
				tabnbr: []int64{1, 2, 3, 4, 5},
			},
		},
		want: ret{
			[]string{"T.data.str", "T.data.tabstr"},
			ErrorGeneric,
		},
	},
	{
		testname: "test 10 - omitempty on struct level with 2 field empty",
		data: &struct {
			data struct {
				str    string
				nbr    int64
				tabstr []string
				tabnbr []int64
			}
		}{
			data: struct {
				str    string
				nbr    int64
				tabstr []string
				tabnbr []int64
			}{
				nbr:    4,
				tabnbr: []int64{1, 2, 3, 4, 5},
			},
		},
		want: ret{
			[]string{"T.data.str", "T.data.tabstr"},
			ErrorGeneric,
		},
	},
}

func TestCheck(t *testing.T) {
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			want := tt.want
			got, err := Check(reflect.ValueOf(tt.data))
			if !reflect.DeepEqual(want, ret{got, err}) {
				t.Errorf("%s:\ngot:\n\t`%v`\nwant:\n\t`%v`", tt.testname, ret{got, err}, want)
			}
		})
	}
}

func TestPanic(t *testing.T) {

}
