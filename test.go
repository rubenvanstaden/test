package test

import (
	"fmt"
	"os"
	"runtime"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/rubenvanstaden/color"
)

// Fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string) {
	if !condition {
        _, file, line, _ := runtime.Caller(1)
        out := fmt.Sprintf(color.InRed("%s:%d\n\n"), filepath.Base(file), line)
        out += fmt.Sprintf(color.InRed("\tmsg: %s\n\n"), msg)
        fmt.Printf(out)
		tb.Fail()
	}
}

// Fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	if err != nil {
        _, file, line, _ := runtime.Caller(1)
        out := fmt.Sprintf(color.InRed("%s:%d\n\n"), filepath.Base(file), line)
        out += fmt.Sprintf(color.InRed("\tcatch: %s\n\n"), err.Error())
        fmt.Printf(out)
		tb.Fail()
	}
}

// Fails the test if "want" is not equal to "got".
func Equals(tb testing.TB, want, got interface{}) {
	if !reflect.DeepEqual(want, got) {
        _, file, line, _ := runtime.Caller(1)
        out := fmt.Sprintf(color.InRed("%s:%d\n\n"), filepath.Base(file), line)
        out += fmt.Sprintf(color.InRed("\twant: %#v\n\n"), want)
        out += fmt.Sprintf(color.InRed("\tgot: %#v\n\n"), got)
        fmt.Printf(out)
		tb.Fail()
	}
}

// Fails the test if the given file doesn't exist.
func Exists(tb testing.TB, filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
        _, file, line, _ := runtime.Caller(1)
        out := fmt.Sprintf(color.InRed("%s:%d\n\n"), filepath.Base(file), line)
        out += fmt.Sprintf(color.InRed("\tno file: %s\n\n"), filename)
        fmt.Printf(out)
		tb.Fail()
	}
}
