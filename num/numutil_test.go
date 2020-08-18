package num

import (
	"github.com/wenj91/util/num"
	"testing"
)

func TestRandomFromTo(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(num.RandomFromTo(1000000000, 9999999999))
	}
}
