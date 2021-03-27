let a = [0, 10, 20, 30, 40]

block: # example 1
   let b = a[2]
   echo b

block: # example 2
   let b = a[2 .. 3]
   echo b

block: # example 3
   let b = a[2 ..< 4]
   echo b

block: # example 4
   let b = a[2 .. ^1]
   echo b
