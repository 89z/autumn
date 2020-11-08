import sequtils
let a0 = ["May", "June"]
# example 1
let s1 = foldl(a0, a & b)
# example 2
proc f(s, s2: string): string = s & s2
let s2 = foldl(a0, f(a, b))
# print
echo s1 == "MayJune" and s2 == "MayJune"
