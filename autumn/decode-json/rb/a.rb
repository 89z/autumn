require 'json'

input = '{"U2": {"Boy": ["Twilight", "I Will Follow"]}}'
m = JSON.parse(input)
out = m['U2']['Boy'][0]
p out == 'Twilight'
