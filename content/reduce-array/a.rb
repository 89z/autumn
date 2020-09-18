a = ['May', 'June']
# example A
sA = a.reduce { |sc, sd| sc + sd }
# example B
f = lambda { |sc, sd| sc + sd }
sB = a.reduce(&f)
# print
puts sA == 'MayJune', sB == 'MayJune'
