package main

import "fmt"

type Handler interface {
	Print(a int)
}

type HandlerFunc func(a int) int

func (f HandlerFunc) Print(a int) {
	v := f(a)
	fmt.Printf("%v\n", v)
}

func main() {
	var f1 HandlerFunc = func(a int) int { return a + 100 }
	f1.Print(100)
}
