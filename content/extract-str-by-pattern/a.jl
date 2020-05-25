s1 = "Wednesday"
s2 = r"e."
# example 1
o1 = match(s2, s1)
s3 = o1.match
println(s3)
# example 2
o2 = eachmatch(s2, s1)
a1 = collect(o2)
println(a1)
