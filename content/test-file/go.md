---
title: Go
tags: [go]
---

{{< r "a.go" >}}

<https://golang.org/pkg/os#FileMode.IsRegular>

Note that for existing files, these are four times slower:

{{< r "b.go" >}}

<https://golang.org/pkg/os#Open>

{{< r "c.go" >}}

<https://golang.org/pkg/os#OpenFile>

## References

- <https://hyperpolyglot.org/c#file-test>
- <https://programming-idioms.org/idiom/144>
- <https://rosettacode.org/wiki/Check_that_file_exists>
