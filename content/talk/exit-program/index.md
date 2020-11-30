---
title: Exit program
---

Rust needs a single line way to exit a program. With Go, you have `Log.Fatal`:

{{< r "a.go" >}}

This works for a `Result`:

{{< r "a.rs" >}}

This works for `Option`:

{{< r "b.rs" >}}

but what if you have both in the same function? Normally we could use this:

{{< r "c.rs" >}}

but it doesnt work for numbers:

{{< r "d.rs" >}}
