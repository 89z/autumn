import strutils
let s1 = "sunday"
# example 1
let b1 = s1.contains("und")
# example 2
let b2 = s1.endsWith("day")
# print
echo [b1, b2]
