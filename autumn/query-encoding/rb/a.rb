require 'cgi'
s = 'west=left&east=right'
m = CGI.parse(s)
p m
