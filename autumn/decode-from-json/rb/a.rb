require 'json'
s = '{"month": 12, "day": 31}'
m = JSON.parse(s)
n = m['day']
p n == 31
