---
title: 'Talk:Two-way conditions'
---

## If statements

PHP accepts any type with conditions:

{{< r "if.php" >}}

Go accepts booleans with conditions:

{{< r "if.go" >}}

but not other types:

{{< r "if_2.go" >}}

To deal with this, we should simply only use booleans when a condition is
expected.

## Ternary operations

PHP offers this:

{{< r "ternary.php" >}}

<https://php.net/operators.comparison#language.operators.comparison.ternary>

Go does not offer this:

> The reason `?:` is absent from Go is that the language's designers had seen
> the operation used too often to create impenetrably complex expressions. The
> `if-else` form, although longer, is unquestionably clearer. A language needs
> only one conditional control flow construct.

<https://golang.org/doc/faq#Does_Go_have_a_ternary_form>

We can continue to use ternary opeations, as long as the condition is a
boolean. For example with Python:

{{< r "ternary.py" >}}

or JavaScript:

{{< r "ternary.js" >}}

## Binary operators

The first argument is a condition. As such, per the previous decision, it must
be a boolean. For example with JavaScript:

{{< r "binary.js" >}}

or PHP:

{{< r "binary.php" >}}

This isnt really useful if we want to return non boolean values. Further, some
languages like PHP coerce the result to boolean:

{{< r "binary_2.php" >}}

or worse, only fail on null:

{{< r "binary_3.php" >}}

<https://php.net/operators.comparison#language.operators.comparison.coalesce>

In order to have two values that can be returned, we would need to use an if
statement or similar. However implementing an if statement means that the
second argument is also a condition, and must be a boolean as well. Better
would be to use boolean type with all arguments. For example with JavaScript:

{{< r "binary_2.js" >}}

Further, some languages like Go only accept boolean values:

{{< r "binary.go" >}}

<https://golang.org/ref/spec#Logical_operators>
