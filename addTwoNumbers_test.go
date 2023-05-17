package main

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	type numbersTestData struct {
		name string
		n1 int
		n2 int
		ans int
	}
	
	tableTest := []numbersTestData{
		{
			name: "1",
			n1: 1,
			n2: 1,
			ans: 2,
		},
		{
			name: "2",
			n1: 2,
			n2: 2,
			ans: 4,
		},
		{
			name: "3",
			n1: 5,
			n2: 4,
			ans: 9,
		},
		{
			name: "4",
			n1: 10,
			n2: 123,
			ans: 133,
		},
		{
			name: "wrong test 1",
			n1: 1,
			n2: 2,
			ans: 100,
		},
	}
	
	for _, i := range(tableTest) {
		t.Run(i.name, func(t *testing.T) {
			want := i.ans
			got := AddTwoNumbers(i.n1, i.n2)

			if want != got {
				t.Errorf("want: %d, got: %d", want, got)
			}
		})
	}
}
