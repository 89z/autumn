a = ['May', 'June']
# example 1
s = a.reduce { |sa, sc| sa + sc }
# example 2
f = lambda { |sa, sc| sa + sc }
s2 = a.reduce(&f)
# print
puts s == 'MayJune', s2 == 'MayJune'
