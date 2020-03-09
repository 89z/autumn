---
title: C++
stars: 134910
---

If you have a Go file like this:

{{< r "a.go" >}}

You get expected output. If you have a C++ file like this:

{{< r "a.cpp" >}}

You get gibberish output. Same with this:

{{< r "a2.cpp" >}}

Same with this:

{{< r "a3.cpp" >}}

This only works with Visual Studio:

{{< r "a4.cpp" >}}

<https://stackoverflow.com/questions/19987448/-/20005850>

This is close, but fails if the size is too small:

{{< r "a5.cpp" >}}

<https://stackoverflow.com/questions/45575863/-/45622802>

## Setup

~~~
x86_64-w64-mingw32-g++ -s -static -std=c++17 a.cpp
~~~

## Stars

<https://github.com/tensorflow/tensorflow>
