import
   os,
   strformat,
   times

let old_n = times.cpuTime()

while true:
   os.sleep 10
   let new_n = cpuTime() - old_n
   stdout.write &"{new_n:.2f}\r"
