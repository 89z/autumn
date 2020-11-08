a = ['May', 'June']
# example 1
s1 = a.reduce(:+)
# example 2
s2 = a.reduce { |s, s2| s + s2 }
# example 3
n3 = a.reduce(0) { |n, s| n + s.length }
# print
puts s1 == 'MayJune', s2 == 'MayJune', n3 == 7
