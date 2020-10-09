f = Proc.new { |n, n0| n + n0 }
# example 1
n1 = f.call(10, 11)
# example 2
a = [10, 11]
n2 = a.reduce(&f)
# print
puts n1 == 21, n2 == 21
