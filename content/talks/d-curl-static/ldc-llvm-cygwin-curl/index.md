---
title: LDC LLVM Cygwin cURL
---

This will not work, as Cygwin cURL uses Cygwin NgHTTP2, which only includes the
shared library:

~~~
/usr/x86_64-w64-mingw32/sys-root/mingw/lib/libnghttp2.dll.a
~~~
