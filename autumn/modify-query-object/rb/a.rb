require 'cgi'
s = 'one=odd&two=even'
m = CGI.parse(s)
puts m
