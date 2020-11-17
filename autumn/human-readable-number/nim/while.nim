import strformat

proc numberFormat(n: float): string =
   var
      n = n
      n2 = 0
   while n > 1000:
      n /= 1000
      n2 += 1
   return fmt"{n:.3f}" & ["", " k", " M", " G"][n2]

let s = numberFormat(9012345678'd)
echo s == "9.012 G"
