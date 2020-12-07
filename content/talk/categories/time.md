+++
title = "Time"
+++

If we do our base with UTC, it will be the same short date as my time zone:

{{< r "a.php" >}}

If we do our base with my time zone, short dates will be different. This is
good, as I will know what zone functions are returning:

{{< r "b.php" >}}

~~~
current date object
   -> date object

date modify
   date object, duration -> date object

number to date object
   date number -> date object

return date number
   -> date number
   date object -> date number
   date string -> date number

return date string
   -> date string
   date number -> date string
   date object -> date string

string to date object
   date string -> date object

return duration number
   date object, date object -> duration number

return duration string
   -> duration string
~~~
