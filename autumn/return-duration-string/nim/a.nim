import
   os,
   strformat,
   times

let old_n = cpuTime()

while true:
   sleep 10
   let new_n = cpuTime() - old_n
   stdout.write &"{new_n:.2f}\r"
