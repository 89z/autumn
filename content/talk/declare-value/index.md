---
title: 'Talk:Declare value'
---

Letter | Type
-------|------
`a`    | Array
`b`    | Boolean
`e`    | Error
`f`    | Function
`m`    | Map
`n`    | Number
`o`    | Object
`r`    | Resource
`s`    | String
`t`    | Set (`t` in `set`, also `t` in `true`)
`y`    | Byte (`y` in `byte`)

## Byte

If we use `a1` for Byte Arrays, it will be confusing when used next to
traditional arrays:

{{< r "byte-arr.go" >}}

If we use `s1` for Byte Array, it will be confusing when used next to
traditions strings:

{{< r "byte-str.go" >}}

<https://golang.org/ref/spec#Numeric_types>

## Error

{{< r "err.go" >}}

<https://golang.org/ref/spec#Errors>

## Object

{{< r "obj.py" >}}

<https://docs.python.org/library/urllib.parse.html#urlparse-result-object>

## Resource

Resource is something that is used by PHP. Other languages use Objects.

{{< r "fopen.php" >}}

<https://php.net/function.fopen>
