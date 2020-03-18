let a1 = ["Sun", "Mon", "Tue"]
# string from front
let s1 = a1[0]
# string from back
let s2 = a1[^1]
# array from front
let a2 = a1[0 .. 1]
# array from back
let a3 = a1[^2 .. ^1]
# print
echo s1, s2, a2, a3
