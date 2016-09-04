package ConsistentHash

import (
	"testing"
)

func check(trueInput, expect int, t *testing.T) {
	if trueInput != expect {
		t.Errorf("expect %d, but got %d\n", trueInput, expect)
	}
}
func TestNew(t *testing.T) {
	x := New()
	if x == nil {
		t.Errorf("new ConsistentHash error")
	}

	check(x.replicas, 10, t)
}
