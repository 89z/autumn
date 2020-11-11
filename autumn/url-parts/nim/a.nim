import uri
let o = "https://example.com/one?two=even".parseUri
# example 1
let s1 = o.hostname
# example 2
let s2 = o.query
# print
echo s1 == "example.com" and s2 == "two=even"
