package main

import (
	"bytes"
	"fmt"
)

func main() {
	padNum := 5
	b := byte(padNum)
	c := []byte{b}
	newC := bytes.Repeat(c, 3)
	fmt.Printf("%T,%v", newC, newC)
}
