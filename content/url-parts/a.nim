import uri
let o = "https://example.com/one?two=even".parseUri
# example A
let sA = o.hostname
# example B
let sB = o.path
# example C
let sC = o.query
# print
echo sA == "example.com" and sB == "/one" and sC == "two=even"
