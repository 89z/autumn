f = lambda { |d, e| d + e }

# example 1
n = f.call(4, 5)
puts n == 9

# example 2
n = [4, 5].reduce(&f)
puts n == 9
