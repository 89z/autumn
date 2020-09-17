proc f(sa, sc: string): string =
   return sa & sc

var a = ["May", "June"]
var s = ""

for sc in a.items:
   s = f(s, sc)

echo s == "MayJune"
