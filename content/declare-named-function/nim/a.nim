# example 1
proc f1(s: string, c: char): bool =
   return s[0] == c
# example 2
proc f2(s, c: string): bool =
   return s[.. 0] == c
# print
echo f1("June", 'J') and f2("June", "J")
