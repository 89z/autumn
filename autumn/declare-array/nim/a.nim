block: # example 1
   var a: seq[int]
   echo a

block: # example 2
   var a: seq[int] = @[10, 11]
   echo a

block: # example 3
   var a = @[10, 11]
   echo a

block: # example 4
   var a = @[
      @[10, 11], @[12, 13]
   ]
   echo a
