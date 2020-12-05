import
   os,
   strformat,
   times

proc format(o: TimeInterval): string =
   let mil_s = fmt"{o.milliseconds:03}"
   let sec_s = fmt"{o.seconds:02}"
   return fmt"{o.minutes} m {sec_s} s {mil_s} ms"

let now_o = now()

while true:
   sleep 10
   let o = between(now_o, now())
   stdout.write format(o), "\r"
