require 'uri'
m = {west: 'left', east: 'right'}
s = URI.encode_www_form(m)
puts s
