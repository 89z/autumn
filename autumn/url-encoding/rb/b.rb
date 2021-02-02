require 'uri'
u = URI::HTTP.build(host: 'ruby-lang.org')
s = u.to_s
puts s
