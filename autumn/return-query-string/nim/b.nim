import uri
let u = parseUri("http://nim-lang.org/docs?month=May&day=Friday")
# example 1
let s1 = u.hostname
# example 2
let s2 = u.query
# print
echo s1 == "nim-lang.org" and s2 == "month=May&day=Friday"
