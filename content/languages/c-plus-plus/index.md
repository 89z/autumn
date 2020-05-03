---
title: C++ language
categories: [language]
---

Regarding booleans, it is not ideal that numbers are printed by default:

{{< r "b.cpp" >}}

but to be fair, other languages have similar issues. Here is Go:

{{< r "a.go" >}}

Here is creating a map:

{{< r "a.cpp" >}}

<https://stackoverflow.com/questions/14070940/-/55278718>

HTTP GET:

<https://github.com/mstorsjo/llvm-mingw/issues/95>

## Setup

~~~
clang++ -s -static -std=c++17 a.cpp
~~~

<https://github.com/mstorsjo/llvm-mingw>

Cygwin Mingw compiler fails with Unicode under "cmd.exe":

<https://sourceware.org/pipermail/cygwin/2020-March/243989.html>
