package main

import "fmt"

//go:generate multisetgen -multiset_typename=IntSet -element_typename=int

func main() {
	m := NewIntSet()
	fmt.Println(m.String())
	fmt.Println(m.Count(3))
	m.Insert(3)
	m.Insert(3)
	m.Insert(3)
	m.Insert(3)
	fmt.Println(m.String())
	fmt.Println(m.Count(3))
	m.Insert(1)
	m.Insert(2)
	m.Insert(5)
	m.Insert(7)
	m.Erase(3)
	m.Erase(5)
	fmt.Println(m.Count(3))
	fmt.Println(m.Count(1))
	fmt.Println(m.Count(2))
	fmt.Println(m.Count(5))
}
