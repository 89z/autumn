# example 1
a = ['south', 'north']
b = a.sort
p b

# example 2
a = ['SOUTH', 'north']
b = a.sort do |s, t|
   s.upcase <=> t.upcase
end
p b
