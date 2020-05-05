s1 = "package"
# example 1
o1 = match(r"p.", s1)
s2 = o1.match
println(s2)
# example 2
o2 = eachmatch(r"a.", s1)
a1 = collect(o2)
println(a1)
# example 3
o3 = match(r"p(..)", s1)
a1 = o3.captures
println(a1)
