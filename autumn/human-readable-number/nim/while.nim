import strformat

proc numberFormat(n: float): string =
   var
      n2 = n
      n3 = 0
   while n2 >= 1e3:
      n2 /= 1e3
      n3 += 1
   return &"{n2:.3f}" & ["", " k", " M", " G"][n3]

let s = numberFormat(9012345678e0)
echo s == "9.012 G"
