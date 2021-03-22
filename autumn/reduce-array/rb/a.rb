a = ['May', 'June']

# example 1
s = a.reduce(:+)
puts s == 'MayJune'

# example 2
s = a.reduce { |s, t| s + t }
puts s == 'MayJune'

# example 3
n = a.reduce(0) { |n, s| n + s.length }
puts n == 7
