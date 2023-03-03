package test

import (
	"fmt"
	"os"
	"testing"
)

func TestSyntax(t *testing.T) {
	test(t)
}

func test(t *testing.T) {
	info, err := os.Stat("../conf/conf.ini")
	fmt.Println(info.Name())
	fmt.Println(info.Size())
	fmt.Println(info.Mode().String())
	fmt.Println(info.ModTime().String())
	fmt.Println(info.IsDir())
	fmt.Println(info.Sys())

	fmt.Println()

	fmt.Printf("%v, %v\n", info, err)
}
