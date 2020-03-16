---
title: UPX
---

Input:

~~~go
package main
func main() {
   println("sunday")
}
~~~

~~~sh
$ upx --best sunday.exe
$ wc -c sunday.exe
619008 sunday.exe

$ upx -9 sunday.exe
$ wc -c sunday.exe
621056 sunday.exe

$ upx sunday.exe
$ wc -c sunday.exe
623616 sunday.exe
~~~

<https://en.wikipedia.org/wiki/UPX>
