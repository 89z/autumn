# example 1
proc f1(n: int, n1: int): bool =
   return n > n1

# example 2
proc f2(n, n2: int): bool =
   return n > n2

# print
let b1 = f1(9, 8)
let b2 = f2(9, 8)
echo b1 and b2
