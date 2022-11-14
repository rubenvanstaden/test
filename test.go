package test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/rubenvanstaden/color"
)

// Fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string) {
	if !condition {
        out := fmt.Sprintf(color.InRed("\nmsg: %s\n"), msg)
		tb.Errorf(out)
	}
}

// Fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	if err != nil {
        out := fmt.Sprintf(color.InRed("\ncatch: %s\n"), err.Error())
		tb.Errorf(out)
	}
}

// Fails the test if "want" is not equal to "got".
func Equals(tb testing.TB, want, got interface{}) {
	if !reflect.DeepEqual(want, got) {
        want := fmt.Sprintf(color.InRed("\nwant: %#v\n"), want)
        got := fmt.Sprintf(color.InRed("\ngot: %#v\n"), got)
        out := want + got
		tb.Errorf(out)
	}
}

// Fails the test if the given file doesn't exist.
func Exists(tb testing.TB, filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
        out := fmt.Sprintf(color.InRed("\nno file: %s\n"), filename)
		tb.Errorf(out)
	}
}
