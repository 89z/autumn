+++
title = "Time"
+++

Reference time will be:

~~~
2020-05-04T03:02:01Z
~~~

We want reference time to be UTC. We want each number to be greater than zero.
We want each number to be different. Avoid using `1` for the month, as in number
form `1` can be January or February. Also this chosen time is good because short
dates will be different in my time zone, depending on if function uses UTC or
local. The timestamp for this is:

{{< r "a.php" >}}

We also need this as an offset from 2020, radix 36 encoded.
