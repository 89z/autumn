require 'securerandom'
# example 1
s1 = SecureRandom.hex
# example 2
s2 = SecureRandom.hex(10)
# print
puts s1, s2
