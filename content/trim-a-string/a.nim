import strutils
# example 1
var s1 = " ab "
var s2 = s1.strip
# example 2
var s3 = "ab\n"
s3.stripLineEnd
# print
echo s2 == "ab", s3 == "ab"
