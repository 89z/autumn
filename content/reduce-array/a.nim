import sequtils
let ab = ["May", "June"]
# example 1
let s = foldl(ab, a & b)
# example 2
proc f(sa, sc: string): string = sa & sc
let s2 = foldl(ab, f(a, b))
# print
echo s == "MayJune" and s2 == "MayJune"
