require 'uri'
s = 'west=left&east=right'
m = URI.decode_www_form(s).to_h
p m
