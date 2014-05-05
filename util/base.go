package util

func NewIntMatrix(n, m int) [][]int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, m)
	}
	return mat
}

func NewFloatMatrix(n, m int) [][]float64 {
	mat := make([][]float64, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]float64, m)
	}
	return mat
}

func IsIntSliceSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func ReverseIntSlice(s []int) []int {
	r := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return r
}

func ReverseRuneSlice(s []rune) []rune {
	r := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return r
}
