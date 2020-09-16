import sequtils
let a_src = ["May", "June"]
let s = foldl(a_src, a & b)
echo s == "MayJune"
