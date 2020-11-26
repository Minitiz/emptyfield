package emptyfield

type datatest struct {
	testname string
	opt      string
	want     bool
}

var test = maketest()

func makeTest() []datatest {
	return []datatest{
		{
			testname: "test 1",
			data:     struct{}{},
			want:     false,
		},
		{},
	}
}
