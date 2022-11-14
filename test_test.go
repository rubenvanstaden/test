package test

import (
	"testing"
	"fmt"
)

func TestUnit_SuccessEquals(t *testing.T) {

	Equals(t, "want", "want")
}

func TestUnit_FailEquals(t *testing.T) {

	Equals(t, "want", "act")
}

func TestUnit_SuccessAssert(t *testing.T) {

	Assert(t, true, "testing message")
}

func TestUnit_FailAssert(t *testing.T) {

	Assert(t, false, "testing message")
}

func TestUnit_FailOk(t *testing.T) {

	Ok(t, fmt.Errorf("Some error"))
}

func TestUnit_FailExists(t *testing.T) {

	Exists(t, "/var/tmp/noexist")
}
