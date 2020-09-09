import strutils
let a = ["May", "June"]
# example 1
let s = a.join(",")
# example 2
let s2 = a.join
# print
echo [s == "May,June", s2 == "MayJune"]
