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

func Test_tagOptions_Contains(t *testing.T) {
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

func Test_Contains(t *testing.T) {
	t.Run("main", func(t *testing.T) {
		o := tagOptions{"field"}
		o.Contains("field")
	})

}
