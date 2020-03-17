---
title: UPX
---

Input:

{{< r "a.go" >}}

Result:

~~~sh
$ upx --best a.exe
$ wc -c a.exe
619008 a.exe

$ upx -9 a.exe
$ wc -c a.exe
621056 a.exe

$ upx a.exe
$ wc -c a.exe
623616 a.exe
~~~

<https://en.wikipedia.org/wiki/UPX>
