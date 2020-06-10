import uri
let s = "https://example.com/one?two=even"
let o = s.parseUri
# example 1
let s1 = o.hostname
# example 2
let s2 = o.path
# example 3
let s3 = o.query
# print
echo s1, s2, s3
