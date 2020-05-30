---
title: 'Talk:Command line arguments'
---

It seems that `CommandLineToArgvW` parses differently than `argv` itself. Take
this program:

{{< r "cmp.c" >}}

If you compile with `cc a.c -lmsvcr100`, you get this:

~~~
C:> c-test.exe "a ""b"" c"
0 c-test.exe
1 a "b" c
0 c-test.exe
1 a "b
2 c
~~~

I did a little searching, and I found this:

<https://devblogs.microsoft.com/oldnewthing/20100917-00?p=12833>

and I noticed, he never uses Quote Quote, but Backslash Quote. So I thought to
retry my example:

~~~
C:> c-test.exe "a \"b\" c"
0 c-test.exe
1 a "b" c
0 c-test.exe
1 a "b" c
~~~

The other characters I had trouble with are `&`, `<`, `>`, `^` and `|`.
Sometimes you can just quote them:

~~~
C:\> c-test.exe "a & b < c > d ^ e | f"
0 c-test.exe
1 a & b < c > d ^ e | f
0 c-test.exe
1 a & b < c > d ^ e | f
~~~

but then other times, you have to Quote and Escape them:

~~~
C:\> c-test.exe "a \" b ^& c ^< d ^> e ^^ f ^| g"
0 c-test.exe
1 a " b & c < d > e ^ f | g
0 c-test.exe
1 a " b & c < d > e ^ f | g
~~~

After some different trials, it seems the best option in my opinion is this.
Never quote `"`, but always escape with `\^`. Then just quote everything else:

<http://daviddeley.com/autohotkey/parameters/parameters.htm#WINCMDRULES>

Example:

~~~
C:\> c-test.exe "a "\^"" b & c < d > e ^ f | g"
0 c-test.exe
1 a " b & c < d > e ^ f | g
0 c-test.exe
1 a " b & c < d > e ^ f | g
~~~

Finally, I wanted to test all the other characters to see if any other problems.
Result are as expected:

~~~
C:\> c-test.exe " !"\^""#$%&'()*+,-./:;<=>?@[\]^_`{|}~"
0 c-test.exe
1  !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~
0 c-test.exe
1  !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~
~~~

Finally, I was curious what would happen if I applied all this to the original
examples:

~~~
C:\> python-2.6.2.amd64\python.exe a.py "a "\^"b\^"" c"
a "b" c

C:\> python-2.5.4.amd64\python.exe a.py "a "\^"b\^"" c"
a "b" c
~~~

and Go:

~~~
C:\> go-test.exe "a "\^"b\^"" c"
a "b" c
~~~

{{< r "main.c" >}}
{{< r "wmain.c" >}}

## Dart

{{< r "a.dart" >}}

## Go

{{< r "a.go" >}}

<https://github.com/golang/go/blob/8f4151ea/src/os/exec_windows.go#L161>

## PHP

{{< r "a.php" >}}

## Python

{{< r "a.py" >}}

## Rust

{{< r "a.rs" >}}

## References

<http://daviddeley.com/autohotkey/parameters/parameters.htm#WINCRULES>
