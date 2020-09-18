# example A
proc fA(nY: int, nZ: int): bool =
   return nY > nZ
# example B
proc fB(nY, nZ: int): bool =
   return nY > nZ
# print
echo fA(9, 8) and fB(9, 8)
