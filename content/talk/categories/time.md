+++
title = "Time"
+++

We want reference time to be UTC. We want each number to be greater than zero.
We want each number to be different. Avoid using `1` for the month, as number
form `1` can be January or February. We also want short dates to be different
in my time zone, depending on if function uses UTC or local. Result is:

{{< r "a.go" >}}

Get second difference from 2020, radix 36 encoded:

{{< r "b.go" >}}

Get millisecond difference from 2020, radix 62 encoded:

{{< r "c.go" >}}
