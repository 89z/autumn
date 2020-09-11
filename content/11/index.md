---
title: Iterate a file
categories: [iterate, input-output]
tags: [d, go, php, python, rust]
date: 2020-07-02
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

## Rust

{{< r "a.rs" >}}

<https://doc.rust-lang.org/std/io/trait.BufRead.html#method.lines>

## References

- <https://hyperpolyglot.org/scripting2#iterate-by-line>
- <https://rosettacode.org/wiki/Read_a_file_line_by_line>
