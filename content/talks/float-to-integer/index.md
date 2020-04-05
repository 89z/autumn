---
title: Float to integer
---

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

{{< r "a.go" >}}

however this is a misrepresentation of the problem. A closer example would be:

{{< r "c.go" >}}

which works fine.
