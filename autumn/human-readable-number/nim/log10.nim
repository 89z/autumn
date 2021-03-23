from math import nil
from strformat import fmt

proc numberFormat(d: float): string =
   let e = math.trunc(math.log10(d)) / 3
   let f = d / math.pow(1e3, e)
   return fmt"{f:.3f}" & ["", " k", " M", " G"][int e]

let s = numberFormat(9012345678.0)
echo s == "9.012 G"
