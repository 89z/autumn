---
title: C++
tags: [draft]
---

**Boolean**:

It is not ideal that numbers are printed by default:

{{< r "b.cpp" >}}

but to be fair, other languages have similar issues. Here is Go:

{{< r "a.go" >}}

**Map**:

{{< r "a.cpp" >}}

<https://stackoverflow.com/questions/14070940/-/55278718>

**Setup**

~~~
clang++ -s -static -std=c++17 a.cpp
~~~

<https://github.com/mstorsjo/llvm-mingw>
