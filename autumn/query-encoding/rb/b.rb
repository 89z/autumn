require 'uri'
s = 'month=May&day=Friday'
m = URI.decode_www_form(s).to_h
puts m
