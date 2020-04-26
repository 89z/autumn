---
title: Red
stars: 4046
---

Shared is slow on first compile, but then fast:

~~~
$ red-13feb20-d513c8f9 -c a.red
...compilation time : 38844 ms
...linking time     : 452 ms
...output file size : 1166848 bytes
...output file      : C:\red\libRedRT.dll

...compilation time : 1014 ms
...linking time     : 63 ms
...output file size : 74752 bytes
...output file      : C:\red\a.exe

$ red-13feb20-d513c8f9 -c a.red
...compilation time : 889 ms
...linking time     : 32 ms
...output file size : 74752 bytes
...output file      : C:\red\a.exe
~~~

Static is always slow:

~~~
$ red-13feb20-d513c8f9 -r a.red
...compilation time : 23540 ms
...linking time     : 156 ms
...output file size : 739328 bytes
...output file      : C:\red\a.exe
~~~

<https://github.com/red/red/issues/3412>

## Stars

<https://github.com/red/red>
