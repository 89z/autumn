require 'securerandom'

# example 1
s = SecureRandom.hex
puts s

# example 2
s = SecureRandom.hex(10)
puts s
