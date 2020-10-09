a = ['May', 'June']
# example 1
s1 = a.reduce(:+)
# example 2
s2 = a.reduce { |s, s2| s + s2 }
# example 3
f = lambda { |s, s3| s + s3 }
s3 = a.reduce(&f)
# example 4
n4 = a.reduce(0) { |n, s| n + s.length }
# print
puts s1 == 'MayJune', s2 == 'MayJune', s3 == 'MayJune', n4 == 7
