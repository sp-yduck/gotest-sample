package main_test

import (
	"fmt"
	"testing"

	"github.com/sp-yduck/gotest-sample"
)

func TestMain(m *testing.M) {
	main.Main()
	fmt.Println("hello main_test !!")
}
