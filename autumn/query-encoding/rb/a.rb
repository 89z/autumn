require 'cgi'
s = 'month=May&day=Friday'
m = CGI.parse(s)
puts m
