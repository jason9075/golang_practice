package main

import "fmt"

type MyData struct {
	foo int
	bar string
}

func main() {
	// Basic init
	data := new(MyData)
	fmt.Println(data.foo)
	fmt.Println(data.bar)
	fmt.Println(data)

	// But the default value is zero. So we write this
	data2 := &MyData{
		foo: 100,
		bar: "hello",
	}
	fmt.Println(data2.foo)
	fmt.Println(data2.bar)
	fmt.Println(data2)

	// but when we init map, slice and channel, we usually use "make"

	//slice
	persons := make([]string, 5, 10) // len = 5, cap = 10
	persons = append(persons, "jason")
	persons = append(persons, "alice")
	fmt.Println(persons)
	fmt.Println(persons[5])

	//map
	colors := make(map[string]int)
	colors["red"] = 10
	colors["green"] = 20
	fmt.Println(colors)
	fmt.Println(colors["red"])

}
