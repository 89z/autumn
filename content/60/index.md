---
title: Convert radix
categories: [number]
tags: [d, go, javascript, nim, php, python, rust]
date: 2020-09-11
---

Some faster bitwise operations could be used, but these functions allow to
easily change the radix. These also account for numeric zero input.

## D

{{< r "a.d" >}}

<https://dlang.org/library/std/conv.html>

## Dart

{{< r "a.dart" >}}

<https://api.dart.dev/dart-core/int/toRadixString.html>

## Go

{{< r "big.go" >}}
{{< r "str.go" >}}

- <https://golang.org/pkg/math/big#Int.SetString>
- <https://golang.org/pkg/math/big#Int.Text>
- <https://golang.org/pkg/strings#IndexRune>

## JavaScript

{{< r "a.js" >}}

<https://developer.mozilla.org/Web/JavaScript/Reference/Global_Objects/String/indexOf>

## Nim

{{< r "a.nim" >}}

<https://nim-lang.org/docs/strutils.html>

## PHP

{{< r "a.php" >}}

<https://php.net/function.strpos>

## Python

{{< r "a.py" >}}

<https://docs.python.org/library/stdtypes.html#str.find>

## Rust

{{< r "a.rs" >}}

<https://doc.rust-lang.org/nightly/std/fmt/index.html>

## References

- <https://hyperpolyglot.org/scripting#radix>
- <https://programming-idioms.org/idiom/142>
- <https://rosettacode.org/wiki/Non-decimal_radices/Convert>
- <https://rosettacode.org/wiki/Non-decimal_radices/Output>
