a = ['May', 'June']
f = lambda { |sa, sc| sa + sc }
# example 1
s = a.reduce(&f)
# example 2
s2 = ''
a.each do |sc|
   s2 = f.call(s2, sc)
end
# print
puts s == 'MayJune', s2 == 'MayJune'
