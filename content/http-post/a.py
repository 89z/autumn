from urllib import request, parse
import re
# GET
o1 = request.urlopen('https://github.com')
s_log = o1.read().decode()
# POST
a_auth = re.findall('data-csrf="true" value="([^"]*)"', s_log)
m_data = {}
m_data['value'] = 'aa540'
m_data['authenticity_token'] = a_auth[0]
y1 = parse.urlencode(m_data).encode()
s_url = 'https://github.com/signup_check/username'
o2 = request.Request(s_url, data=y1, method='POST')
request.urlopen(o2)
