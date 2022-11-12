package test

import (
	"testing"
	// "fmt"
)

func TestUnit(t *testing.T) {

	Assert(t, false, "testing message")
	// Ok(t, fmt.Errorf("Some error"))
	// Equals(t, "want", "act")
	// Exists(t, "/var/tmp/noexist")
}
