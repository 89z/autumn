---
title: Byte array
---

## Go

Its important to note that a byte slice is *an array*, not a string:

> A slice is a descriptor for a contiguous segment of an underlying **array**
> and provides access to a numbered sequence of elements from that **array**.

<https://golang.org/ref/spec#Slice_types>

This can be easily demonstrated:

{{< r "a.go" >}}

So as shown above, when using Byte Slices, they shall be declared with `a1` or
similar identifiers, not `s1` or similar. When printing Byte Slices, we could
use one of these:

{{< r "b.go" >}}
{{< r "c.go" >}}

but we are adding an extra line just so that we can use a function that is not
guaranteed to exist:

> Println is useful for bootstrapping and debugging; it is not guaranteed to
> stay in the language.

<https://golang.org/pkg/builtin#println>

This is the better option:

{{< r "d.go" >}}

## Python

Python has 2 builtin libraries for HTTP:

- <https://docs.python.org/library/http.client.html>
- <https://docs.python.org/library/urllib.request.html>

However both libraries return byte arrays:

{{< r "a.py" >}}
{{< r "b.py" >}}

Running a process has a similar problem:

{{< r "c.py" >}}

but with a process, you can open in text mode:

{{< r "d.py" >}}

None of the HTTP methods have a text option:

- {{< a `https://docs.python.org/library/http.client.html#
   http.client.HTTPConnection.getresponse` >}}
- {{< a `https://docs.python.org/library/http.client.html#
   http.client.HTTPConnection.request` >}}
- {{< a `https://docs.python.org/library/http.client.html#
   http.client.HTTPConnection` >}}
- {{< a `https://docs.python.org/library/urllib.request.html#
   urllib.request.urlopen` >}}

Workaround:

{{< r "e.py" >}}
