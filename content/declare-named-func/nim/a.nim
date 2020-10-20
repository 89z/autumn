# example 1
proc f1(n: int, n1: int): bool =
   return n > n1
# example 2
proc f2(n, n2: int): bool =
   return n > n2
# print
echo f1(9, 8) and f2(9, 8)
