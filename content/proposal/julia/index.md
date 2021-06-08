+++
title = "Julia"
tags = [ "proposal" ]
date = 2021-06-08
+++

## Speed of ERROR: LoadError

If I have a file like this:

~~~jl
println("Sunday")
~~~

Julia works fine:

~~~
$ time julia a.jl
Sunday

real    0m0.170s
~~~

With a file like this:

~~~
a()
~~~

The error process seems to be very slow:

~~~
$ time julia b.jl
ERROR: LoadError: UndefVarError: a not defined
Stacktrace:
 [1] top-level scope at D:\Desktop\etc\b.jl:1
 [2] include(::Module, ::String) at .\Base.jl:377
 [3] exec_options(::Base.JLOptions) at .\client.jl:288
 [4] _start() at .\client.jl:484
in expression starting at D:\Desktop\etc\b.jl:1

real    0m0.990s
~~~

The `--compile` flag helps significantly:

~~~
$ time julia --compile=min b.jl
ERROR: LoadError: UndefVarError: a not defined
Stacktrace:
 [1] top-level scope at D:\Desktop\etc\b.jl:1
 [2] include(::Module, ::String) at .\Base.jl:377
 [3] exec_options(::Base.JLOptions) at .\client.jl:288
 [4] _start() at .\client.jl:484
in expression starting at D:\Desktop\etc\b.jl:1

real    0m0.290s
~~~

but still seems to be quite slow compared to other options, like Python:

~~~
$ cat a.py
a()

$ time python a.py
Traceback (most recent call last):
  File "a.py", line 1, in <module>
    a()
NameError: name 'a' is not defined

real    0m0.060s
~~~

PHP:

~~~
$ cat a.php
<?php
a();

$ time php a.php
Fatal error: Uncaught Error: Call to undefined function a() in C:\a.php:2
Stack trace:
#0 {main}
  thrown in C:\a.php on line 2

real    0m0.060s
~~~

Dart:

~~~
$ cat a.dart
main() {
   a();
}

$ time dart a.dart
a.dart:2:4: Error: Method not found: 'a'.
   a();
   ^

real    0m0.160s
~~~

- <https://github.com/JuliaLang/julia/issues/17285>
- <https://lwn.net/Articles/856819>
