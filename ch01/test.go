package main

import "fmt"

type People struct {
	name string
	age int
}

func main() {
	p1 := &People{}
	p1.name = "yipeng"
	p1.age = 31
	fmt.Println(p1)
}
