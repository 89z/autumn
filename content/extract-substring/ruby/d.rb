s = 'March'
# example 1
s1 = s[... 2]
# example 2
s2 = s[2 ...]
# example 3
s3 = s[2 ... 4]
# example 4
s4 = s.slice(... 2)
# example 5
s5 = s.slice(2 ...)
# example 6
s6 = s.slice(2 ... 4)
# print
puts s1 == 'Ma', s2 == 'rch', s3 == 'rc', s4 == 'Ma', s5 == 'rch', s6 == 'rc'
