package Week06

import (
	"testing"
	"time"
)

func TestSlidingWindowCounter(t *testing.T) {
	rn := NewRollingNumber(2, 1)

	rn.IncrSuccess()
	rn.IncrSuccess()
	rn.IncrFail()
	time.Sleep(time.Second * 1)
	rn.IncrSuccess()
	rn.IncrFail()
	time.Sleep(time.Second * 1)
	rn.IncrSuccess()
	time.Sleep(time.Second * 1)

	if rn.GetRollingSum().success == 2 && rn.GetRollingSum().fail == 1 {
		t.Log("满足预期")
	} else {
		t.Error("不满足预期")
	}
}
