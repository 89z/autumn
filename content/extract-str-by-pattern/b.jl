s1 = "Wednesday"
s2 = r"e(..)"
# match captures
o = match(s2, s1)
a = o.captures
println(a)
# eachmatch collect
o = eachmatch(s2, s1)
a = collect(o)
println(a)
