import math, strformat

proc numberFormat(n: float): string =
   let n2 = int log10(n) / 3
   let n = n / pow(1e3, float n2)
   return fmt"{n:.3f}" & ["", " k", " M", " G"][n2]

let s = numberFormat(9012345678e0)
echo s == "9.012 G"
