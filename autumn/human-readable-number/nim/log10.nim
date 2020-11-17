import math, strformat

proc numberFormat(n: float): string =
   let n2 = int log10(n) / 3
   let n = n / pow(1000, float n2)
   return fmt"{n:.3f}" & ["", "  k", " M", " G"][n2]

let s = numberFormat(9012345678'd)
echo s == "9.012 G"
