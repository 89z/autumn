require 'uri'
o = URI('https://example.com/one?two=even')
puts s1 == 'example.com', s2 == 'two=even'
s2 = o.query
