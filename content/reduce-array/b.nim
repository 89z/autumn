let a_src = ["May", "June"]
var s = ""

proc f(s_acc, s_cur: string): string =
   return s_acc & s_cur

for s_cur in a_src.items:
   s = f(s, s_cur)

echo s == "MayJune"
