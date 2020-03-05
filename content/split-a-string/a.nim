import strutils
let s1 = "sun mon tue"
let s2 = """sun mon tue
wed thu fri"""
# example 1
let a1 = s1.split(" ")
# example 2
let a2 = s1.split(maxsplit = 1)
# example 3
let a3 = s2.split
# example 4
let a4 = s2.splitLines
# print
echo a1, "\n", a2, "\n", a3, "\n", a4
