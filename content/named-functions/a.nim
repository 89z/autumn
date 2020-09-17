# example 1
proc f(n: int, n2: int): bool =
   return n > n2
# example 2
proc f2(n, n2: int): bool =
   return n > n2
# print
let b = f(9, 8)
let b2 = f2(9, 8)
echo b and b2
