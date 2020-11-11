require 'uri'
o = URI('https://example.com/one?two=even')
# example 1
s1 = o.host
# example 2
s2 = o.query
# print
puts s1 == 'example.com', s2 == 'two=even'
