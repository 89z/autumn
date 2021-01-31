require 'uri'
s = 'one=odd&two=even'
m = URI.decode_www_form(s).to_h
puts m
