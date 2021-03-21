import strtabs

block: # example 1
   let m = newStringTable()
   echo m

block: # example 2
   let m = {"one": "even", "two": "odd"}.newStringTable
   echo m
