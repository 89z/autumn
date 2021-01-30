import uri
let u = parseUri("https://example.com/one?two=even")
# example 1
let s1 = u.hostname
# example 2
let s2 = u.query
# print
echo s1 == "example.com" and s2 == "two=even"
