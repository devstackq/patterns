package open_closed

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/*
open to extend -
close to modify
interface  case = polymophysm;
new struct -> new realize -> extend
no modify - called function; only call; object; mongo{} || postgres{} , etc object


Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification.
    be able to override a struct
    Use the Strategy Design Pattern
*/
/*function example
Because countWords now works with any string and we can create new functions by composing it
(we can extend its behaviour) and without modifying its code (it is closed for modification).
*/

func countWords(s string) int {
	return len(strings.Split(s, " "))
}

func countWordsUrl() int {
	file, _ := ioutil.ReadFile("/path_to_file")
	return countWords(string(file))
}

// func countWordsFile() int {
// 	resp, _ := http.Get("http://example.com/")
// 	return countWords(string(resp.Body()))
// }

// be able to override a struct

type A struct{ year int }

func (a A) Greet() {
	fmt.Println("Hello GolangUK", a.year)
}

type B struct {
	A
}

func (b B) Greet() {
	fmt.Println("Welcome to GolangUK", b.year)
}

// pattern strategy
type Calculer interface {
	Execute(int, int) int
}
type (
	Add      struct{}
	Minus    struct{}
	Devide   struct{}
	Multiple struct{}
)

func (add Add) Execute(a, b int) int {
	return a + b
}

func (m Minus) Execute(a, b int) int {
	return a - b
}

// can extend; new struct - own logic
func (d Devide) Execute(a, b int) int {
	return a / b
}

func (d Multiple) Execute(a, b int) int {
	return a * b
}

type Calcul struct {
	c Calculer
}

func (c Calcul) Calculate(a, b int) int {
	return c.c.Execute(a, b)
}

func Test() {
	az := Add{}
	r := az.Execute(120, 31)

	add := Calcul{c: Add{}}
	minus := Calcul{c: Minus{}}
	devide := Calcul{c: Devide{}}

	a := add.Calculate(120, 31)
	b := minus.Calculate(41, 11)
	c := devide.Calculate(a, b)
	log.Println(c, "res, pattern strategy", r)
}
