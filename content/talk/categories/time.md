---
title: Time
---

If we do our base with UTC, it will be the same short date as my time zone:

{{< r "a.php" >}}

If we do our base with my time zone, short dates will be different. This is
good, as I will know what zone functions are returning:

{{< r "b.php" >}}

- current date object
- date addition
- number to date object
- return date number
   - from null
   - from string
   - from object
- return date string
   - from null
   - from number
   - from object
- return duration number
- return duration object
- return duration string
- string to date object
