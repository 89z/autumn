import sequtils
let a_mon = ["May", "June"]
let s_mon = foldl(a_mon, a & b)
echo s_mon == "MayJune"
