---
title: Go bytes.Buffer
---

What is difference between `bytes.Buffer` and `strings.Builder`?

~~~
String_Builder/1Write_NoGrow-4    20000000  60.4 ns/op   48 B/op  1 allocs/op
String_Builder/3Write_NoGrow-4    10000000   230 ns/op  336 B/op  3 allocs/op
String_Builder/3Write_Grow-4      20000000   102 ns/op  112 B/op  1 allocs/op
String_ByteBuffer/1Write_NoGrow-4 10000000   125 ns/op  160 B/op  2 allocs/op
String_ByteBuffer/3Write_NoGrow-4  5000000   339 ns/op  400 B/op  3 allocs/op
String_ByteBuffer/3Write_Grow-4    5000000   316 ns/op  336 B/op  3 allocs/op
~~~

<https://go-review.googlesource.com/c/go/+/96980>
