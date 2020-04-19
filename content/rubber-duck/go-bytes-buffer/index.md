---
title: Go bytes.Buffer
---

## Go

Its important to note that a byte slice is *an array*, not a string:

> A slice is a descriptor for a contiguous segment of an underlying **array**
> and provides access to a numbered sequence of elements from that **array**.

<https://golang.org/ref/spec#Slice_types>

When using Byte Slices, they shall be declared with `a1` or similar
identifiers, not `s1` or similar. When printing Byte Slices, we could use one
of these:

{{< r "c.go" >}}

but we are adding an extra line just so that we can use a function that is not
guaranteed to exist:

> Println is useful for bootstrapping and debugging; it is not guaranteed to
> stay in the language.

<https://golang.org/pkg/builtin#println>

What is difference between `bytes.Buffer` and `strings.Builder`?

## Python

Python has 2 builtin libraries for HTTP:

- <https://docs.python.org/library/http.client.html>
- <https://docs.python.org/library/urllib.request.html>

However both libraries return byte arrays:

{{< r "b.py" >}}

Running a process has a similar problem, but with a process, you can open in
text mode. None of the HTTP methods have a text option:

- {{< a `https://docs.python.org/library/http.client.html#
   http.client.HTTPConnection.getresponse` >}}
- {{< a `https://docs.python.org/library/http.client.html#
   http.client.HTTPConnection.request` >}}
- {{< a `https://docs.python.org/library/http.client.html#
   http.client.HTTPConnection` >}}
- {{< a `https://docs.python.org/library/urllib.request.html#
   urllib.request.urlopen` >}}

Workarounds:

{{< r "e.py" >}}
{{< r "f.py" >}}

Here are some topics:

- byte types
- bytes to console
- string to console
- string types

before we go down this road again, we need to confront the fact that some
languages have no bytes type, namely PHP. So as long as that is the case, any
byte topics need to live in Rubber Duck Debug.
