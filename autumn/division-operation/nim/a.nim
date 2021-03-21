let
   f = 7.5
   i = 7

block: # example 1
   let n = f / 2
   echo n == 3.75

block: # example 2
   let n = int(f / 2)
   echo n == 3

block: # example 3
   let n = int(i / 2)
   echo n == 3

block: # example 4
   let n = i / 2
   echo n == 3.5
