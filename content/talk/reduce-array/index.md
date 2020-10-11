---
title: Reduce array
---

We can approach this three ways. Using the reduce function:

{{< r "a.php" >}}

Functional loop:

{{< r "b.go" >}}
{{< r "b.php" >}}

Inline loop:

{{< r "c.go" >}}
{{< r "c.php" >}}

The reduce function can be used in cases where you dont need or dont care about
the current index, and just want some short code.

The inline loop can be used in cases where reduce function is not available
(Go), or in cases wheres reduce callback does not offer current index (PHP).

The functional loop can be used when you need a loop, but you want to reuse the
inner code.
