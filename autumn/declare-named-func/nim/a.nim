# example 1
proc addf(n: int, n1: int): int =
   return n + n1

# example 2
proc subf(n, n2: int): int =
   return n - n2

# print
let n1 = addf(8, 1)
let n2 = subf(8, 1)
echo n1 == 9 and n2 == 7
