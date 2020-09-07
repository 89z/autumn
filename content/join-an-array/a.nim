import strutils
let a = ["Sunday", "Monday"]
# example 1
let s = a.join(",")
# example 2
let s2 = a.join
# print
echo [s == "Sunday,Monday", s2 == "SundayMonday"]
