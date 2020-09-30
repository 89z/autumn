# example 1
f1 = Proc.new { |n, n1| n > n1 }
# example 2
f2 = proc { |n, n2| n > n2 }
# example 3
f3 = lambda { |n, n3| n > n3 }
# exmaple 4
f4 = lambda do |n, n4|
   return n > n4
end
# print
b1 = f1.call(9, 8)
b2 = f2.call(9, 8)
b3 = f3.call(9, 8)
b4 = f4.call(9, 8)
puts b1, b2, b3, b4
