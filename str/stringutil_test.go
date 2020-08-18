package str

import (
	"testing"
)

func TestToString(t *testing.T) {
	st := String(1, 2, 3, 4, 5, 6)
	t.Log(st)
}