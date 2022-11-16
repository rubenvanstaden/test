package test

import (
	"testing"
	"fmt"
)

func TestUnit_SuccessEquals(t *testing.T) {
	Equals(t, "want", "want")
}

func TestUnit_SuccessOk(t *testing.T) {

	Ok(t, nil)
}

func TestUnit_SuccessExists(t *testing.T) {

	// Exists(t, "/var/tmp/noexist")
}

func TestUnit_FailEquals(t *testing.T) {

	Equals(t, "want", "act")
}

func TestUnit_FailOk(t *testing.T) {

	Ok(t, fmt.Errorf("Some error"))
}

func TestUnit_FailExists(t *testing.T) {

	Exists(t, "/var/tmp/noexist")
}
