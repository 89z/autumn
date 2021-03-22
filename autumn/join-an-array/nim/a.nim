import strutils
let a = ["May", "June"]

block: # example 1
   let s = a.join(",")
   echo s == "May,June"

block: # example 2
   let s = a.join
   echo s == "MayJune"
