def f(sa, sc)
   return sa + sc
end

a = ['May', 'June']
s = ''

a.each do |sc|
   s = f(s, sc)
end

puts s == 'MayJune'
