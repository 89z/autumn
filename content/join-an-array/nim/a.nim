import strutils
let a = ["May", "June"]
# example 1
let s1 = a.join(",")
# example 2
let s2 = a.join
# print
echo s1 == "May,June" and s2 == "MayJune"
