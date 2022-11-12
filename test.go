package test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	Black  = "\033[38m"
	White  = "\033[39m"
)

// Fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string) {
	if !condition {
		printPosition()
		fmt.Printf(Red+"\tmsg: %s\n\n"+Reset, msg)
		tb.FailNow()
	}
}

// Fails the test if an err is not nil.
func Ok(tb testing.TB, err error) {
	if err != nil {
		printPosition()
		fmt.Printf(Red+"\tcatch error: %s\n\n"+Reset, err.Error())
		tb.FailNow()
	}
}

// Fails the test if exp is not equal to act.
func Equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		printPosition()
		fmt.Printf(Red+"\twant: %#v\n\n"+Reset, exp)
		fmt.Printf(Red+"\tgot: %#v\n\n"+Reset, exp)
		tb.FailNow()
	}
}

func Exists(tb testing.TB, filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		printPosition()
		fmt.Printf(Red+"\tno file: %s\n\n"+Reset, filename)
		tb.FailNow()
	}
}

func printPosition() {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf(Red+"%s:%d\n\n"+Reset, filepath.Base(file), line)
}
