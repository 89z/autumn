---
title: 'Talk:Convert to integer'
---

## Float

These functions:

{{< r "a.php" >}}

all return floats rather than integers:

~~~
float(2)
float(2)
float(1)
~~~

with PHP this is fine:

{{< r "b.php" >}}

but with other languages like Go it can be a problem:

{{< r "a.go" >}}

workaround:

{{< r "b.go" >}}

however this is a misrepresentation of the problem. A closer example would be:

{{< r "c.go" >}}

which works fine. Even this is a misrepresentation, as we want rounding rather
than truncation. A closer example would be:

{{< r "d.go" >}}

## String

Test           | Ops/sec
---------------|-----------
`Number('10')` | 68,962,875
`+('10')`      | 36,446,186

<https://jsperf.com/number-vs-parseint-vs-plus>
