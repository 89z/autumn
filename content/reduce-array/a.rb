a = ['May', 'June']
# example 1
s1 = a.reduce { |s, s1| s + s1 }
# example 2
f = lambda { |s, s2| s + s2 }
s2 = a.reduce(&f)
# print
puts s1 == 'MayJune', s2 == 'MayJune'
