import strutils
var s1 = "sun mon tue\n"
# example 1
var s2 = s1.strip
# example 2
s1.stripLineEnd
# print
echo [s1, s2]
