import strutils
let s1 = "sun mon tue"
# example 1
let n1 = s1.find(" ")
# example 2
let n2 = s1.rfind(" ")
# print
echo [n1, n2]
