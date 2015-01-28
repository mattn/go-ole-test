package foo

import (
	"github.com/mattn/go-ole"
)

func Foo() {
	ole.CoInitialize(0)
}
