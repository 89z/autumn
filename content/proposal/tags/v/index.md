---
title: V
tags: [draft]
---

## field designator  languagetm_gmtoff' does not refer to any field

Trying to build my own version creates this error:

~~~
$ TMPDIR=tmp PATH=../llvm-mingw/bin:$PATH make
cd tmp/vc && git clean -xf && git pull --quiet
cd /var/tmp/tcc && git clean -xf && git pull --quiet
cc  -std=gnu11 -w -o v tmp/vc/v.c  -lm
tmp/vc/v.c:9234:38: error: field designator  languagetm_gmtoff' does not refer to any
      field in type  languagestruct tm'
                                    .tm_gmtoff = 0};
                                     ^
~~~

## Compiler emits no warnings or errors

If I have a file like this, everything is fine:

~~~v
fn main() {
   mut m2 := { languageSun': 10, 'Mon': 11}
   m2[ languageTue'] = 12
   println("Sunday")
}
~~~

If I remove `mut`, it will cause a problem:

~~~v
fn main() {
   m2 := { languageSun': 10, 'Mon': 11}
   m2[ languageTue'] = 12
   println("Sunday")
}
~~~

but instead of printing a warning or error, the compiler just exits. Does V
have a way to emit messages when a problem happens?
Workaround is Self-compiled binary.

## Program generated after compiling on Windows display nothing

Workaround:

~~~
PATH=llvm-mingw/bin v_windows/v
~~~

<https://github.com/vlang/v/issues/3052>
