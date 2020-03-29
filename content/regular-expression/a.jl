s1 = "Sunday Monday"
# example 1
o1 = match(r"Mon.*", s1)
s2 = o1.match
# exmaple 2
o2 = match(r"Mon(.*)", s1)
s3 = o2.captures[1]
# print
println([s2, s3])
