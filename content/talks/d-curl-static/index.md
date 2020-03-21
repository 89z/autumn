---
title: D cURL static
---

First a Hello World with DMD and statically linked C library:

{{< r "hello/hello.c" >}}
{{< r "hello/hello.d" >}}
{{< r "hello/hello_test.d" >}}
{{< r "hello/hello.sh" >}}

Result:

~~~
dmd
OPTLINK (R) for Win32  Release 8.00.17
Copyright (C) Digital Mars 1989-2013  All rights reserved.
http://www.digitalmars.com/ctg/optlink.html
hello.lib
 Error 43: Not a Valid Library File
~~~

DMD only offers the shared library:

~~~
windows\bin\libcurl.dll
windows\bin64\libcurl.dll
~~~

Next we can try GDC. This is not viable as last release is 2016:

<https://gdcproject.org/downloads>

Next we can try LDC. LDC only offers the shared library:

~~~
bin\libcurl.dll
lib\libcurl.dll
~~~

Next we need to post or link to issues about missing static library:

- <https://github.com/ldc-developers/ldc/issues/3376>
- <https://issues.dlang.org/show_bug.cgi?id=15841>
- <https://issues.dlang.org/show_bug.cgi?id=20690>

Next we will build our own static library:

{{< r "curl/curl.sh" >}}

Next we can try to use static library with `std.net.curl` and `etc.c.curl`.

{{< r "curl/etc_c_curl.d" >}}

## References

- <https://dlang.org/blog/2017/10/25/dmd-windows-and-c>
- <https://github.com/ldc-developers/ldc/issues/3260>
- <https://lld.llvm.org>
