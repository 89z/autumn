require 'securerandom'

# example 1
s = SecureRandom.base64
puts s

# example 2
s = SecureRandom.base64(10)
puts s
