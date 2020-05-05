---
title: 'Talk:Slice a string'
---

We want to return begin character(s), as syntax can be different:

{{< r "beg.go" >}}

We want to return middle character(s), as Go can have unexpected results:

{{< r "mid.go" >}}

We also want to return end character(s), as Go can have unexpected results:

{{< r "end.go" >}}

We also want to return a single character, as Go can have unexpected results:

{{< r "one.go" >}}

We also want to return multiple characters, as JavaScript can have unexpected
results:

{{< r "two.js" >}}

Make the begin a single character:

{{< r "beg-one.js" >}}

Result:
