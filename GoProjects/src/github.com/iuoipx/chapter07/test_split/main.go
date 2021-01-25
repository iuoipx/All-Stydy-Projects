package main

import (
	"fmt"

	"github.com/iuoipx/chapter07/split_string"
)

func main() {
	ret := split_string.Split("abcabsdnd", "b")
	fmt.Printf("%#v\n", ret)
	ret1 := split_string.Split("bcb", "b")
	fmt.Printf("%#v\n", ret1)
	ret2 := split_string.Split("bbb", "b")
	fmt.Printf("%#v\n", ret2)
}
