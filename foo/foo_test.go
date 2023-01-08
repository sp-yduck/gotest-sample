package foo_test

import (
	"fmt"
	"testing"

	"github.com/sp-yduck/gotest-sample/foo"
)

func TestFoo(t *testing.T) {
	foo.Foo()
	fmt.Println("hello foo test !!")
}
