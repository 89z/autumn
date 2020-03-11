---
title: C++
stars: 134910
---

Here is writing to the console:

{{< r "a.cpp" >}}

Here is creating a map:

{{< r "a2.cpp" >}}

<https://stackoverflow.com/questions/14070940/-/55278718>

Regarding booleans, it is not ideal that numbers are printed by default:

{{< r "a3.cpp" >}}

but to be fair, other languages have similar issues. Here is Go:

{{< r "a.go" >}}

## Setup

~~~
clang++ -s -static -std=c++17 a.cpp
~~~

<https://github.com/mstorsjo/llvm-mingw>

Cygwin Mingw compiler fails with Unicode under "cmd.exe":

<https://sourceware.org/pipermail/cygwin/2020-March/243989.html>

## Stars

<https://github.com/tensorflow/tensorflow>
