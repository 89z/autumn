from os import nil
from strformat import fmt
from times import TimeInterval

proc format(o: TimeInterval): string =
   let mil_s = fmt"{o.milliseconds:03}"
   let sec_s = fmt"{o.seconds:02}"
   return fmt"{o.minutes} m {sec_s} s {mil_s} ms"

let now_o = times.now()

while true:
   os.sleep 10
   let o = times.between(now_o, times.now())
   stdout.write format(o), "\r"
