require 'uri'
u = URI('http://ruby-lang.org?west=left&east=right')
s = u.query
puts s
