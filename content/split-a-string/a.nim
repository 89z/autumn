import strutils
# example 1
let a1 = "Sun Mon".split(" ")
# example 2
let a2 = "Sun Mon Tue".split(maxsplit = 1)
# example 3
let a3 = "Sun Mon".split
# example 4
let a4 = """Sunday
Monday""".splitLines
# print
echo a1, a2, a3, a4
