require 'securerandom'
# example 1
s1 = SecureRandom.base64
# example 2
s2 = SecureRandom.base64(10)
# print
puts s1, s2
