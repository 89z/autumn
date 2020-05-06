# example 1
o1 = match(r"^Su", s1)
println(o1 != nothing)
# example 3
o2 = match(r"un.", s1)
println(o2 != nothing)
# example 5
o3 = match(r"ay$", s1)
println(o3 != nothing)
