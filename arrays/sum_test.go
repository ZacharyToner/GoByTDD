package arrays

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if want != got {
		t.Errorf("get %d but wanted %d, given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{7, 8, 9}, []int{0, 9})
	want := []int{6, 24, 9}

	if want != got {
		t.Errorf("get %v but wanted %v", got, want)
	}
}
