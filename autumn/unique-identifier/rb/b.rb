require 'securerandom'
# example 1
s1 = SecureRandom.alphanumeric
# example 2
s2 = SecureRandom.alphanumeric(10)
# print
puts s1, s2
