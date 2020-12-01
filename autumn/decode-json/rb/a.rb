require 'json'
in_s = '{"U2": {"Boy": ["Twilight", "I Will Follow"]}}'
m = JSON.parse(in_s)
out_s = m['U2']['Boy'][0]
puts out_s == 'Twilight'
