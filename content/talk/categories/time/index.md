---
title: 'Category talk:Time'
---

Some languages like Go cannot convert directly from number to date string:

{{< r "a.go" >}}

Same with JavaScript:

{{< r "a.js" >}}

PHP can:

{{< r "a.php" >}}

and PHP can also convert from object:

{{< r "b.php" >}}

we could approach like this:

- get date number
- get date object
- get date string

but those pages could be too large, as could be multiple input and output
options. Better would be this:

- number to object
- number to string
- object to number
- object to string
- string to number
- string to object

So in this case, we can split like this:

- number to string
- number to object
