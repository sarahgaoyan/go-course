package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Eat() {
	fmt.Printf("human %v can eat\n", h)
}

type Student struct {
	Human
	Grade int
	Class int
}

func (s *Student) Study() {
	fmt.Printf("student %v can study\n", s)
}

/* interface */
type Animal interface {
	GetColor() string
	GetCatalog() string
}

type Cat struct {
	Color   string
	Catalog string
}

func (c *Cat) GetColor() string {
	fmt.Println("I am white")
	return c.Color
}

func (c *Cat) GetCatalog() string {
	fmt.Println("I am cat")
	return c.Catalog
}

type Dog struct {
	Color   string
	Catalog string
}

func (d *Dog) GetColor() string {
	fmt.Println("I am yellow")
	return d.Color
}

func (d *Dog) GetCatalog() string {
	fmt.Println("I am dog")
	return d.Catalog
}

func main() {
	h := Human{"a", 10}
	h.Eat()
	s := Student{h, 1, 1}
	s.Eat()
	s.Study()

	fmt.Println("==== interface ===")
	var a Animal
	a = &Cat{"white", "cat"}
	a.GetColor()
	a.GetCatalog()

	a = &Dog{"yellow", "dog"}
	a.GetColor()
	a.GetCatalog()
}
