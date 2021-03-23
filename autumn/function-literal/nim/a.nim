block: # example 1
   let f = proc (d: int, e: int): int =
      return d + e
   echo f(4, 5) == 9

block: # example 2
   let f = proc (d, e: int): int =
      return d + e
   echo f(4, 5) == 9

block: # example 3
   let f = proc (d, e: int): int = d + e
   echo f(4, 5) == 9
