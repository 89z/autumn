import sequtils
let c = ["May", "June"]

block: # example 1
   let s = foldl(c, a & b)
   echo s == "MayJune"

block: # example 2
   proc f(s, t: string): string = s & t
   let s = foldl(c, f(a, b))
   echo s == "MayJune"
