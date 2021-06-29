+++
title = "Go"
tags = [ "go" ]
+++

55.90 ns/op:

{{< r "a.go" >}}

<https://golang.org/pkg/strings#SplitN>

102.6 ns/op:

{{< r "b.go" >}}

<https://golang.org/pkg/strings#Fields>

136.8 ns/op:

{{< r "c.go" >}}

<https://golang.org/pkg/strings#Split>

250.0 ns/op:

{{< r "d.go" >}}

<https://golang.org/pkg/strings#FieldsFunc>

682.9 ns/op:

{{< r "e.go" >}}

<https://golang.org/pkg/bufio#Scanner.Split>

## References

- <https://hyperpolyglot.org/c#split>
- <https://programming-idioms.org/idiom/49>
- <https://rosettacode.org/wiki/Tokenize_a_string>
