package werr

import (
	"errors"
	"fmt"
	"runtime"
	"strings"

	"golang.org/x/xerrors"
)

func source() error {
	return errors.New("base error")
}

func caller1() error {
	return Wrap(source())
}

func caller2() error {
	return Wrap(caller1())
}

var (
	_ error             = (*werr)(nil)
	_ fmt.Formatter     = (*werr)(nil)
	_ xerrors.Formatter = (*werr)(nil)
	_ xerrors.Wrapper   = (*werr)(nil)
)

var ErrBase = errors.New("base error")

func ExampleWrap() {
	base := source()
	fmt.Printf("%v\n", base)

	level1 := caller1()
	fmt.Printf("level 1: %v\n", level1)
	fmt.Println(stripPath(fmt.Sprintf("level 1 detail: %+v\n", level1)))

	level2 := caller2()
	fmt.Printf("level 2: %v\n", level2)
	fmt.Println(stripPath(fmt.Sprintf("level 2 detail: %+v", level2)))

	// Output:
	// base error
	// level 1: base error
	// level 1 detail: base error:
	//     github.com/iand/werr.caller1
	//         {path}/werr/err_test.go:17
	//
	// level 2: base error
	// level 2 detail: base error:
	//     github.com/iand/werr.caller1
	//         {path}/werr/err_test.go:17
	//     github.com/iand/werr.caller2
	//         {path}/werr/err_test.go:21
}

func stripPath(s string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return s
	}

	if idx := strings.LastIndex(filename, "/werr/err_test.go"); idx != -1 {
		path := filename[:idx]
		s = strings.ReplaceAll(s, path, "{path}")
	}

	return s
}
