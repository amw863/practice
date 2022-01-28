package recursion

import "testing"

func TestNum(t *testing.T) {
	if num(10) == 10 {
		t.Logf("success")
	}
}
