package main

import (
	"fmt"
	"strconv"
)

type qint int
func (q qint) String() string {
	return fmt.Sprint(`"`, strconv.Itoa(int(q)), `"`)
}

type RString interface {
	rstring()
}

type qstring string
func (q qstring) String() string {
	return fmt.Sprint(`"`, string(q), `"`)
}
func (q qstring) rstring() {}

type env string
func (e env) String() string {
	return fmt.Sprint("`echo $", string(e), "`")
}
func (e env) rstring() {}

func main() {
	fmt.Println(qint(10))
	foo(qstring("123"))
	foo(env("PORT"))

	fmt.Println(`foo\nbar`)
	fmt.Println("foo\nbar")
	fmt.Println("foo\\nbar")
}

func foo(s RString) {
	fmt.Println(s)
}
