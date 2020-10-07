s = 'June'
# example 1
s1 = s[... 2]
# example 2
s2 = s[1 ...]
# example 3
s3 = s[1 ... 3]
# example 4
s4 = s.slice(... 2)
# example 5
s5 = s.slice(1 ...)
# example 6
s6 = s.slice(1 ... 3)
# print
puts s1 == 'Ju', s2 == 'une', s3 == 'un', s4 == 'Ju', s5 == 'une', s6 == 'un'
