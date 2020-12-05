from math import nil
from os import nil
from strformat import fmt
from times import nil

proc format(n: float): string =
   let mil_n = int(math.mod(n, 1) * 1000)
   let sec_n = int(n)
   let min_n = sec_n div 60
   return fmt"{min_n} m {sec_n:02} s {mil_n:03} ms"

let n = times.cpuTime()

while true:
   os.sleep 10
   let s = format(times.cpuTime() - n)
   stdout.write s, "\r"
