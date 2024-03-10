package main

import (
	"fmt"
	"sort"
	"strings"
)

// Задание 12
// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.


type StringSet map[string]struct{}


func NewStringSet(v []string) StringSet {
	set := make(StringSet)
	for _, s := range v {
		set[s] = struct{}{}
	}
	return set
}


func (s *StringSet) Has(elem string) bool {
	_, ok := (*s)[elem]
	return ok
}


func (s *StringSet) Add(elem string) {
	(*s)[elem] = struct{}{}
}


func (s *StringSet) Delete(elem string) {
	delete(*s, elem)
}


func (s *StringSet) Clear() {
	*s = make(StringSet)
}


func (s *StringSet) String() string {
	elements := make([]string, 0, len(*s))
	for element := range *s {
		elements = append(elements, element)
	}
	
	sort.Strings(elements)
	return fmt.Sprintf("[%s]", strings.Join(elements, ", "))
}

func main() {
	v := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewStringSet(v)
	fmt.Println(set.String())

	set.Add("gopher")
	set.Add("gopher") 
	fmt.Println(set.String())

	fmt.Println(set.Has("tree"))
	set.Delete("tree")
	set.Delete("tree") 
	fmt.Println(set.Has("tree"))

	fmt.Println(set.String())
	set.Clear()
	fmt.Println(set.String())
}
