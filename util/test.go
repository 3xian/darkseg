package util

import (
	"fmt"
	"testing"
)

func Expect(t *testing.T, result interface{}, expect interface{}) {
	if fmt.Sprint(result) != fmt.Sprint(expect) {
		t.Errorf("\nResult: %v\nExpect: %v\n", result, expect)
	}
}
