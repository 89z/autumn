require 'json'
s = '{"month": 12, "day": 31}'
m = JSON.parse(s)
puts m['day'] == 31
