---
title: Rust
stars: 43062
---

Rust does not have Map literals:

<https://github.com/rust-lang/rfcs/issues/542>

but it does have `struct`:

{{< r "a.rs" >}}

<https://doc.rust-lang.org/book/ch05-02-example-structs.html>

and `match`:

<https://doc.rust-lang.org/book/match.html>

Go `struct` allows missing fields:

{{< r "a.go" >}}

However Rust does not:

{{< r "b.rs" >}}

Instead requiring something like this:

{{< r "c.rs" >}}

<https://github.com/rust-lang/rfcs/issues/2875>

## HTTP

<https://github.com/sbstp/attohttpc>

## Setup

<https://forge.rust-lang.org/infra/other-installation-methods>

## Stars

<https://github.com/rust-lang/rust>
