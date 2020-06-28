---
title: Iterate a file
categories: [iterate, input-output]
tags: [d, go, php, python]
---

## D

{{< r "a.d" >}}

<https://dlang.org/library/std/stdio/file.by_line.html>

## Go

{{< r "a.go" >}}

## PHP

{{< r "a.php" >}}

`feof` is triggered after `fgets`. Further, `trim` expects a string, so we
cannot use it until we are certain that we have a string.

<https://php.net/function.feof>

## Python

{{< r "a.py" >}}

<https://docs.python.org/reference/compound_stmts.html#for>

## References

- <https://hyperpolyglot.org/scripting2#iterate-by-line>
- <https://rosettacode.org/wiki/Read_a_file_line_by_line>
