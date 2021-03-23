import strformat

proc numberFormat(d: float): string =
   var
      e = d
      f = 0
   while e >= 1e3:
      e /= 1e3
      f += 1
   return &"{e:.3f}" & ["", " k", " M", " G"][f]

let s = numberFormat(9012345678e0)
echo s == "9.012 G"
