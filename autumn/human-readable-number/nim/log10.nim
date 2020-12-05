from math import nil
from strformat import fmt

proc numberFormat(n: float): string =
   let n2 = math.trunc(math.log10(n)) / 3
   let n3 = n / math.pow(1e3, n2)
   return fmt"{n3:.3f}" & ["", " k", " M", " G"][int n2]

let s = numberFormat(9012345678e0)
echo s == "9.012 G"
