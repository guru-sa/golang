package main

import (
	"flag"
	"fmt"

	bytecounter "github.com/ys-sa/go/tutorial/interface/bytecounter"
	tempconv "github.com/ys-sa/go/tutorial/interface/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the ")

func main() {
	var c bytecounter.ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	flag.Parse()
	fmt.Println(*temp)
}
