require 'securerandom'

# example 1
s = SecureRandom.urlsafe_base64
puts s

# example 2
s = SecureRandom.urlsafe_base64(10)
puts s
