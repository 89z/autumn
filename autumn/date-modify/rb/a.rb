require 'date/delta'
s = "10 min, 3 s"
delta = Date::Delta.parse(s).in_secs
t = Time.now + delta
