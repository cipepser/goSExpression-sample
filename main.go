package main

import (
	"fmt"
	"strings"
)

type Tree struct {
	Parent *Tree
	Child  []*Tree
	Value  string
}

func NewTree() *Tree {
	t := &Tree{}
	t.Parent = t
	return t
}

func Parse(s string) (*Tree, error) {
	if s[0] != '(' {
		return nil, fmt.Errorf("first letter must be '(', have '%s'", string(s[0]))
	}
	t := NewTree()

	t = t.Insert(s[1:])

	return t, nil
}

func (t *Tree) Insert(s string) *Tree {
	if s == "" {
		return t
	}
	s = strings.TrimSpace(s)

	switch s[0] {
	case '(':
		n := &Tree{Parent: t}
		n = n.Insert(s[1:])
		t.Child = append(t.Child, n)
	case ')':
		t.Parent = t.Parent.Insert(s[1:])

	default:
		for i := 0; i < len(s); i++ {

			if s[i] == '(' || s[i] == ')' {
				t.Value = s[:i]
				t = t.Insert(s[i:])
				break
			}

		}
	}
	return t
}

func (t *Tree) Walk() {
	fmt.Println("my parent is ", t.Parent.Value)
	fmt.Println(t.Value)
	fmt.Println("-------")

	for _, c := range t.Child {
		c.Walk()
	}
}

func main() {
	s := `(ROOT (S (PP (NP (JJ Natural) (NN language) (NN processing)) )))`
	t, _ := Parse(s)
	t.Walk()
}
