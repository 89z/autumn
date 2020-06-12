let s1 = "Sunday"
# example 1
let s2 = s1[s1.high - 1 .. s1.high]
# example 2
let s3 = s1[^2 .. ^1]
# print
echo [s2 == "ay", s3 == "ay"]
