package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, tab("NAME", "ENABLED"))
	fmt.Fprintln(w, tab("hello", "true"))
	fmt.Fprintln(w, tab("goodbye", "true"))
	fmt.Fprintln(w, tab("very long farewell", "false"))
}

func tab(ss ...string) string {
	return strings.Join(ss, "\t")
}
