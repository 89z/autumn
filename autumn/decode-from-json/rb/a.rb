require 'json'

src = '{"U2": {"Boy": ["Twilight", "I Will Follow"]}}'
m = JSON.parse(src)
dst = m['U2']['Boy'][0]
p dst == 'Twilight'
