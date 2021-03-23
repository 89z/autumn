from os import nil
from strformat import fmt
from times import TimeInterval

proc format(t: TimeInterval): string =
   let mil = fmt"{t.milliseconds:03}"
   let sec = fmt"{t.seconds:02}"
   return fmt"{t.minutes} m {sec} s {mil} ms"

let then = times.now()

while true:
   os.sleep 10
   let now = times.between(then, times.now())
   stdout.write format(now), "\r"
