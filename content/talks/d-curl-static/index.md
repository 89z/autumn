---
title: D cURL static
---

The first choice for this task is `std.net.curl`. However D does not provide a
static library:

<https://issues.dlang.org/show_bug.cgi?id=15841>

We can build cURL ourself:

~~~
./buildconf
./configure --host=x86_64-w64-mingw32 --disable-shared
make LDFLAGS=-all-static
~~~

## References

- <https://dlang.org/blog/2017/10/25/dmd-windows-and-c>
- <https://github.com/ldc-developers/ldc/issues/3260>
