package util

import "testing"

func TestIsIntSliceSame(t *testing.T) {
	Expect(t, IsIntSliceSame([]int{}, []int{}), true)
	Expect(t, IsIntSliceSame([]int{1}, []int{1}), true)
	Expect(t, IsIntSliceSame([]int{1, 11}, []int{1, 11}), true)

	Expect(t, IsIntSliceSame([]int{11}, []int{1, 11}), false)
	Expect(t, IsIntSliceSame([]int{1, 11}, []int{1}), false)
	Expect(t, IsIntSliceSame([]int{}, []int{1}), false)
	Expect(t, IsIntSliceSame([]int{1, 11}, []int{}), false)
}

func TestReverseIntSlice(t *testing.T) {
	Expect(t, ReverseIntSlice([]int{1, 2, 3}), []int{3, 2, 1})
}
