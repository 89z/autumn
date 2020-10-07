s = 'May'
# example 1
s1 = s.slice(1)
# example 2
s2 = s.slice(1, 1)
# example 3
s3 = s.slice(1 .. 1)
# example 4
s4 = s.slice(1 ... 2)
# print
puts s1 == 'a', s2 == 'a', s3 == 'a', s4 == 'a'
