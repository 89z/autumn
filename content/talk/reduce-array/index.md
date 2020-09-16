---
title: 'Talk:Reduce array'
date: 2020-09-14
---

We can approach this three ways. Using the reduce function:

{{< r "a.php" >}}

Functional loop:

{{< r "b.go" >}}
{{< r "b.php" >}}

Inline loop:

{{< r "c.go" >}}
{{< r "c.php" >}}

The reduce function can be used in cases where you dont need or dont care about
the current index, and just want some short code.

The inline loop can be used in cases where reduce function is not available
(Go), or in cases wheres reduce callback does not offer current index (PHP).

The functional loop can be used when you need a loop, but you want to reuse the
inner code.

## Nim

You can use `foldl` like this:

~~~nim
import sequtils
let a_src = ["May", "June"]
let s = foldl(a_src, a & b)
echo s == "MayJune"
~~~

but it seem you cannot supply an actual function:

~~~nim
import sequtils
let a_src = ["May", "June"]

proc f(s_acc, s_cur: string): string =
   return s_acc & s_cur

# Error: type mismatch: got
# <proc (s_acc: string, s_cur: string): string{.noSideEffect, gcsafe, locks: 0.}>
# but expected 'string'
foldl(a_src, f)
~~~

This is possible with other languages, like C++:

~~~c
#include <iostream>
#include <numeric>
#include <vector>

std::string f(std::string s_acc, std::string s_cur) {
   return s_acc + s_cur;
}

int main() {
   std::vector<std::string> a = {"May", "June"};
   auto s = std::accumulate(a.begin(), a.end(), std::string(), f);
   std::cout << (s == "MayJune") << std::endl;
}
~~~

D:

~~~d
import std.algorithm, std.stdio;

void main() {
   auto a = ["May", "June"];
   auto f = (string s_acc, string s_cur) => s_acc ~ s_cur;
   auto s = a.reduce!(f);
   writeln(s == "MayJune");
}
~~~

Python:

~~~py
import functools
a = ['May', 'June']
f = lambda s_acc, s_cur: s_acc + s_cur
s = functools.reduce(f, a)
print(s == 'MayJune')
~~~
