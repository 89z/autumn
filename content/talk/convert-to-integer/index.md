---
title: 'Talk:Convert to integer'
date: 2020-09-14
---

## Float

Some functions return floats rather than integers. With PHP this is fine, but
with other languages like Go it can be a problem. However this is a
misrepresentation of the problem. Even this is a misrepresentation, as we want
rounding rather than truncation.

## String

Test           | Ops/sec
---------------|-----------
`Number('10')` | 68,962,875
`+('10')`      | 36,446,186

<https://jsperf.com/number-vs-parseint-vs-plus>
