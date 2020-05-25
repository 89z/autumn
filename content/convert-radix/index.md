---
title: Convert radix
categories: [number]
tags: [go, javascript, php, python]
---

Some faster bitwise operations could be used, but these functions allow to
easily change the radix. These also account for numeric zero input.

## Go

{{< r "a.go" >}}
{{< r "b.go" >}}

- <https://golang.org/pkg/math/big#Int.SetString>
- <https://golang.org/pkg/math/big#Int.Text>
- <https://golang.org/pkg/strings#IndexRune>

## JavaScript

{{< r "a.js" >}}

<https://developer.mozilla.org/Web/JavaScript/Reference/Global_Objects/String/indexOf>

## PHP

{{< r "a.php" >}}

<https://php.net/function.strpos>

## Python

{{< r "a.py" >}}

<https://docs.python.org/library/stdtypes.html#str.find>

## References

- <https://hyperpolyglot.org/scripting#radix>
- <https://rosettacode.org/wiki/Non-decimal_radices/Convert>
