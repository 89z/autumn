# example 1
f1 = Proc.new { |n, n2| n > n2 }
# example 2
f2 = proc { |n, n2| n > n2 }
# example 3
f3 = lambda { |n, n2| n > n2 }
# exmaple 4
f4 = lambda do |n, n2|
   return n > n2
end
# print
b1 = f1.call(9, 8)
b2 = f2.call(9, 8)
b3 = f3.call(9, 8)
b4 = f4.call(9, 8)
puts b1, b2, b3, b4
