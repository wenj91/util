package str

import (
	"fmt"
	"strings"
)

// String obj to string
func String(args ...interface{}) string {
	return strings.TrimSuffix(fmt.Sprintln(args...), "\n")
}