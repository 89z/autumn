require 'uri'
u = URI('http://ruby-lang.org?month=May&day=Friday')
s = u.query
puts s
