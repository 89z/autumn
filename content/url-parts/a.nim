import uri
let o = "https://example.com/one?two=even".parseUri
# example 1
let s = o.hostname
# example 2
let s2 = o.path
# example 3
let s3 = o.query
# print
echo [s == "example.com", s2 == "/one", s3 == "two=even"]
