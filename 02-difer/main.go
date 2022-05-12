package main

// copy example from https://go.dev/doc/effective_go#examples
import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
}

// result:

// entering: b
// in b
// entering: a
// in a
// leaving: a
// leaving: b
