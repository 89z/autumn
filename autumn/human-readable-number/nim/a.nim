import strformat

proc numberFormat(n: float): string =
   var
      n = n
      n2 = 0
   while n > 1024:
      n /= 1024
      n2 += 1
   return fmt"{n:.3f}" & ["", " k", " M", " G"][n2]

let s = numberFormat(1264)
echo s == "1.234 k"
