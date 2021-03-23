from math import nil
from os import nil
from strformat import fmt
from times import nil

proc format(n: float): string =
   let mil = int(math.mod(n, 1) * 1000)
   let sec = int(n)
   let mins = sec div 60
   return fmt"{mins} m {sec:02} s {mil:03} ms"

let n = times.cpuTime()

while true:
   os.sleep 10
   let s = format(times.cpuTime() - n)
   stdout.write s, "\r"
