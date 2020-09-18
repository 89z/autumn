import sequtils
let aZ = ["May", "June"]
# example A
let sA = foldl(aZ, a & b)
# example B
proc f(sY, sZ: string): string = sY & sZ
let sB = foldl(aZ, f(a, b))
# print
echo sA == "MayJune" and sB == "MayJune"
