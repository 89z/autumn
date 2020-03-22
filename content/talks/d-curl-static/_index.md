---
title: D cURL static
---

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

## References

- <https://dlang.org/blog/2017/10/25/dmd-windows-and-c>
- <https://github.com/ldc-developers/ldc/issues/3260>
- <https://lld.llvm.org>
- <https://stackoverflow.com/questions/2734719>
- <https://support.microsoft.com/help/2977003>
