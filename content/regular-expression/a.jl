s1 = "2020-03-29"
# example 1
o1 = match(r"-..", s1)
s2 = o1.match
println(s2)
# example 2
o2 = match(r"-(..)-(..)", s1)
a1 = o2.captures
println(a1)
# example 3
o3 = eachmatch(r"-..", s1)
a2 = collect(o3)
println(a2)
