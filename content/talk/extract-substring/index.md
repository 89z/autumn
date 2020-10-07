---
title: 'Talk:Extract substring'
---

Points of ambiguity:

1. How many characters are returned?
2. Is second argument length, inclusive or exclusive?

First question is solved by showing the result, example with JavaScript:

{{< r "a.js" >}}

For second question, we could try some examples. This fails as it could be
length or exclusive:

{{< r "b.js" >}}

This fails as it could be length of exclusive:

{{< r "c.js" >}}

This is good:

{{< r "d.js" >}}

This is good:

{{< r "e.js" >}}
