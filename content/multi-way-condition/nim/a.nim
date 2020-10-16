var s = "minute"
var n: int

case s
of "hour":
   n = 23
of "minute", "second":
   n = 59
else:
   n = -1

echo n == 59
