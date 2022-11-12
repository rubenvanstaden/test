package test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/rubenvanstaden/color"
)

// Fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string) {
	if !condition {
		printPosition()
		fmt.Printf(color.InRed("\tmsg: %s\n\n"), msg)
		tb.FailNow()
	}
}

// Fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	if err != nil {
		printPosition()
		fmt.Printf(color.InRed("\tcatch: %s\n\n"), err.Error())
		tb.FailNow()
	}
}

// Fails the test if "want" is not equal to "got".
func Equals(tb testing.TB, want, got interface{}) {
	if !reflect.DeepEqual(want, got) {
		printPosition()
		fmt.Printf(color.InRed("\twant: %#v\n\n"), want)
		fmt.Printf(color.InRed("\tgot: %#v\n\n"), want)
		tb.FailNow()
	}
}

// Fails the test if the given file doesn't exist.
func Exists(tb testing.TB, filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		printPosition()
		fmt.Printf(color.InRed("\tno file: %s\n\n"), filename)
		tb.FailNow()
	}
}

func printPosition() {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf(color.InRed("%s:%d\n\n"), filepath.Base(file), line)
}
