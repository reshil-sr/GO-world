package sum

import "testing"

func TestSum(t *testing.T) {
	s := Sums(1, 2, 3, 4, 5)
	if s != 15 {
		t.Errorf("The sum of 1-5 is %d but got %d", 15, s)
	}
}

func TestSumMul(t *testing.T) {
	tt := []struct {
		inp []int
		opt int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{nil, 0},
		{[]int{1, -1}, 0},
	}
	for _, inpts := range tt {
		s := Sums(inpts.inp...)
		if s != inpts.opt {
			t.Errorf("The sum of %v is %d but got %d", inpts.inp, inpts.opt, s)
		}
	}
}

func TestSumMulAsSubTest(t *testing.T) {
	tt := []struct {
		name string
		inp  []int
		opt  int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"No numbers", nil, 0},
		{"one & minus one", []int{1, -1}, 0},
	}
	for _, inpts := range tt {
		t.Run(inpts.name, func(t *testing.T) {
			s := Sums(inpts.inp...)
			if s != inpts.opt {
				t.Errorf("The sum of %v is %d but got %d", inpts.inp, inpts.opt, s)
			}
		})
	}
}
