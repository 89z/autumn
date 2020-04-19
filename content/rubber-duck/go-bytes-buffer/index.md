---
title: Go bytes.Buffer
---

## Go

Its important to note that a byte slice is *an array*, not a string:

> A slice is a descriptor for a contiguous segment of an underlying **array**
> and provides access to a numbered sequence of elements from that **array**.

<https://golang.org/ref/spec#Slice_types>

When using Byte Slices, they shall be declared with `a1` or similar
identifiers, not `s1` or similar.

What is difference between `bytes.Buffer` and `strings.Builder`?

## Python

Python has 2 builtin libraries for HTTP:

- <https://docs.python.org/library/http.client.html>
- <https://docs.python.org/library/urllib.request.html>

However both libraries return byte arrays. Running a process has a similar
problem, but with a process, you can open in text mode. None of the HTTP
methods have a text option:

- {{< a `https://docs.python.org/library/http.client.html#
   http.client.HTTPConnection.getresponse` >}}
- {{< a `https://docs.python.org/library/http.client.html#
   http.client.HTTPConnection.request` >}}
- {{< a `https://docs.python.org/library/http.client.html#
   http.client.HTTPConnection` >}}
- {{< a `https://docs.python.org/library/urllib.request.html#
   urllib.request.urlopen` >}}
