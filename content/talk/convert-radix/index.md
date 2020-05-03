---
title: 'Talk:Convert radix'
---

If you percent encode everything from U+0021 to U+007E, then you have:

~~~
0-9
a-z
A-Z
-
.
_
~~~

which is 65 characters. Omit the `.` and that is 64.

## C

{{< r "glibc.c" >}}

<https://github.com/lattera/glibc/blob/master/stdlib/l64a.c>

{{< r "musl.c" >}}

<https://github.com/bminor/musl/blob/master/src/misc/a64l.c>

{{< r "newlib.c" >}}

<https://github.com/mirror/newlib-cygwin/blob/master/newlib/libc/stdlib/l64a.c>
