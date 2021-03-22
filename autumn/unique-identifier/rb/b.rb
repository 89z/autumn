require 'securerandom'

# example 1
s = SecureRandom.alphanumeric
puts s

# example 2
s = SecureRandom.alphanumeric(10)
puts s
