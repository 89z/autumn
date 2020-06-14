---
title: Red
tags: [reject]
---

Shared is slow on first compile, but then fast:

~~~
red -c a.red
~~~

Static is always slow:

~~~
PS C:\red> .\red-14jun20-a34a784f8.exe -r a.red

-=== Red Compiler 0.6.4 ===-

Compiling C:\red\a.red ...
Compiling compression library...
...compilation time : 735 ms

Target: MSDOS

Compiling to native code...
...compilation time : 27665 ms
~~~

<https://github.com/red/red/issues/3412>
