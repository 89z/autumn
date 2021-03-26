block: # natural float
   let n = 7.0 / 2
   echo n == 3.5

block: # natural float
   let n = 7 / 2
   echo n == 3.5

block: # force int
   let n = int(7 / 2)
   echo n == 3
