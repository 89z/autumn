let s = "May"
# example 1
let s2 = s[1 .. 1]
# example 2
let s3 = s[1 .. ^1]
# example 3
let s4 = s[1 .. s.high]
# print
echo s2 == "a" and s3 == "ay" and s4 == "ay"
