# example 1
proc f(d: int, e: int): int =
   return d + e

# example 2
proc g(d, e: int): int =
   return d + e

# print
echo f(4, 5) == 9
echo g(4, 5) == 9
