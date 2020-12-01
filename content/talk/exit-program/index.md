---
title: Exit program
---

Rust needs a single line way to exit a program. With Go, you have `Log.Fatal`:

{{< r "a.go" >}}

This works for a `Result`:

{{< r "a.rs" >}}

This works for `Option`:

{{< r "b.rs" >}}

but what if you have both in the same function? Need something that works for
numbers as well. Go has an option:

{{< r "b.go" >}}

Can we use `Err(format!())` for everything? Yes we can:

{{< r "f.rs" >}}

We can even do better:

{{< r "g.rs" >}}

but then what about traditional errors? Yep:

{{< r "h.rs" >}}

Now lets combine:

{{< r "i.rs" >}}

If we dont like the long line, we can do this:

{{< r "j.rs" >}}
