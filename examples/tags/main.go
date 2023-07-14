package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Conf struct {
	Foo string `conf:"foo"`
	Num int `conf:"num"`
}

func quotes(s string) string {
	return fmt.Sprint(`"`, s, `"`)
}

func main() {
	c := Conf{
		Foo: "bar",
		Num: 123,
	}

	sb := &strings.Builder{}
	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)

	sb.WriteString("conf(\n")

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup("conf")
		if !ok {
			continue
		}

		value := fmt.Sprintf("%v", v.Field(i).Interface())

		sb.WriteString(fmt.Sprint("\t", tag, "=", quotes(value), "\n"))
	}

	sb.WriteString(")\n")

	fmt.Println(sb.String())
}