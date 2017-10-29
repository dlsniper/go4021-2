package main

import "some"
import "some/dark"

func main() {
	a := &some.Thing{}
	dark.Hello(a)
}
