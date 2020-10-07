---
title: Red
tags: [proposal]
date: 2020-10-05
---

Shared is slow on first compile, but then fast:

~~~
red -c a.red
~~~

Static is always slow:

~~~
red-14jun20-a34a784f8.exe -r a.red
~~~

<https://github.com/red/red/issues/3412>