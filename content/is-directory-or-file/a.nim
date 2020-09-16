import os
let s = "index.md"
# example 1
let b = s.existsFile
# example 2
let b2 = not s.existsDir
# print
echo b and b2
