package bilist

import (
	// "container/list"

	"fmt"
	"strings"
)

type Tree struct {
	Parent, Child *Tree
	Values        []string
}

func NewTree() *Tree {
	t := new(Tree)
	// t.Child = t
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
	fmt.Println("-----------")
	fmt.Println("t: ", &t)
	fmt.Println("start: ", s)
	fmt.Println("s[0]: ", string(s[0]))

	switch s[0] {
	case '(':
		fmt.Println("t: ", t)
		fmt.Println("t.Child: ", t.Child)
		fmt.Println("t.Parent: ", t.Parent)

		n := &Tree{}
		n.Parent = t
		n = n.Insert(s[1:])
		t.Child = n
		// t.Child = n.Insert(s[1:])
		// t.Child.Parent = t
		// fmt.Println(t.Child.Parent)

		// t.Child = t.Child.Insert(s[1:])

	case ')':
		if t.Parent == t {
			return t
		}
		fmt.Println(":::::", t)
		t.Parent.Insert(s[1:])

	default:
		t.Values = append(t.Values, string(s[0]))
		fmt.Println(t.Values)

		t = t.Insert(s[1:])
	}

	return t
}

func (t *Tree) Walk() {
	for _, v := range t.Values {
		fmt.Print(v, " ")
	}
	if t == t.Child || t == nil {
		return
	}
	fmt.Println("")
	t.Child.Walk()
}

func bilist() {
	// 	s := `(ROOT
	//     (S
	//       (PP
	//         (NP
	//           (JJ Natural)
	//           (NN language)
	//           (NN processing)
	//         )
	//         (IN From)
	//         (NP
	//           (NNP Wikipedia)
	//         )
	//       )
	//       (, ,)
	//       (NP
	//         (NP
	//           (DT the)
	//           (JJ free)
	//           (NN encyclopedia)
	//           (JJ Natural)
	//           (NN language)
	//           (NN processing)
	//         )
	//       )
	//     )
	//   )
	// `
	// s := `(ROOT (S (PP (NP (JJ Natural) (NN language) (NN processing)) )))`
	// s := `(a(b(c(d(e)(f)(g)))))`
	// s := `(a(b)(c)(d))`
	s := `(a(b))`
	// s := `(a)`

	t, _ := Parse(s)
	// fmt.Println(t)
	// t := NewTree()
	t.Walk()

	// l := list.New()
	// l.PushFront("a")
	// l.PushFront("b")
	// fmt.Println(l.PushFront("a"))
	// fmt.Println(l)
	// l.PushFront(v)
}
