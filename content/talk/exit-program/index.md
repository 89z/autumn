+++
title = "Exit program"
+++

Rust needs a single line way to exit a program. With Go, you have `Log.Fatal`.
What if you have both in the same function? Need something that works for
numbers as well. Can we use `Err(format!())` for everything? Yes we can:

{{< r "f.rs" >}}

What about traditional errors? Now lets combine:

{{< r "i.rs" >}}

If we dont like the long line, we can do this:

{{< r "j.rs" >}}
