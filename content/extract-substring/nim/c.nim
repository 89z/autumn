let s = "May"
# example 1
let s1 = s[1 ..< 2]
# example 2
let s2 = s[1 ..< ^0]
# example 3
let s3 = s[1 ..< s.len]
# print
echo s1 == "a" and s2 == "ay" and s3 == "ay"
