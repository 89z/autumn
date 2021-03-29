import strutils
let a = ["West", "East"]

block: # example 1
   let s = a.join(",")
   echo s == "West,East"

block: # example 2
   let s = a.join
   echo s == "WestEast"
