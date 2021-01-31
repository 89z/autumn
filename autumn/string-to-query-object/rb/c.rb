require 'uri'
o = URI('https://example.com/one?two=even')
s1 = o.host
puts s1 == 'example.com'
