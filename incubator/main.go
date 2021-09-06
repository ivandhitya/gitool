package main

import "fmt"

type Model struct {
	Name *string
	Age  *int
}

func main() {
	name := "ivan"
	model := Model{Name: &name}
	fmt.Printf("%v\n", *model.Name)
}
