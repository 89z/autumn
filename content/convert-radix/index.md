---
title: Convert radix
topics: [number]
langs: [javascript, python]
---

## JavaScript

{{< r "a.js" >}}

## Python

{{< r "a.py" >}}

## Notes

If you percent encode everything from U+0021 to U+007E, then you have:

~~~
0-9
a-z
A-Z
-
.
_
~~~

which is 65 characters. Omit the `.` and that is 64.

<https://codegolf.stackexchange.com/questions/536>
