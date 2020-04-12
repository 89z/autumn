---
title: Process
---

So initially I was going to use something from this list:

<https://pubs.opengroup.org/onlinepubs/9699919799/utilities>

but this list is not portable like the name suggests:

~~~
C:\> sleep
'sleep' is not recognized as an internal or external command,
operable program or batch file.
~~~

plus many of the tools are old and crappy, with crap ways of proposing issues or
changes. I would rather promote tools that I actually use and care about. So I
start with this search:

<https://github.com/search?q=stars:%3E10000+size:%3C10000>

and I came up with this:

<https://github.com/ggreer/the_silver_searcher>

~~~sh
# example 1
ag
# exmaple 2
ag -V
~~~

both run under standard output, so no problem.
