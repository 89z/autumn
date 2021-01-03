require 'securerandom'
# example 1
s1 = SecureRandom.urlsafe_base64
# example 2
s2 = SecureRandom.urlsafe_base64(10)
# print
puts s1, s2
