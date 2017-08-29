# gose

gose is a simple s-expression parser. We assume the followings.

* literal has ONLY `string`.
* parse into n-ary tree.



### Install

```
# go get github.com/cipepser/goSExpression-sample/gose

```

### Usage

You can parse s-expression to use `Parse()` and print the Tree by `Walk()`.

```

s := `(a(b(d(i)(j))(e)(f(k)))(c(g)(h(l(n))(m))))`
t, err := gose.Parse(s)

if err != nil {
  log.Fatal(err)
}

t.Walk()

```

## References
* [ROSETTACODE.ORG](https://rosettacode.org/wiki/S-Expressions#Go)