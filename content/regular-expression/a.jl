s1 = "2019-12-31"
# example 1
o1 = match(r"\d", s1)
b1 = o1 != nothing
println(b1)
# example 2
o2 = match(r"-..", s1)
s2 = o2.match
println(s2)
# example 3
o3 = match(r"-(..)-(..)", s1)
a1 = o3.captures
println(a1)
# example 4
o4 = eachmatch(r"-..", s1)
a2 = collect(o4)
println(a2)
