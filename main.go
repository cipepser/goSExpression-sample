package main

import (
	"fmt"

	"./gose"
)

func main() {
	s := `(a(b(d(i)(j))(e)(f(k)))(c(g)(h(l(n))(m))))`
	// s := `(a(b)(c))`
	t, _ := gose.Parse(s)
	fmt.Println(t.Child[0])
	fmt.Println(t.Child[1])
	t.Walk()
}
