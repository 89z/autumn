# example 1
let f = proc (n: int, n2: int): bool =
   return n > n2

# example 2
let f2 = proc (n, n2: int): bool =
   return n > n2

# example 3
let f3 = proc (n, n2: int): bool = n > n2

# print
let b = f(9, 8)
let b2 = f2(9, 8)
let b3 = f3(9, 8)
echo b and b2 and b3
