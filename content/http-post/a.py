from http import cookiejar
from urllib import request, parse
import re
o1 = request.build_opener(request.HTTPCookieProcessor())
# GET
s_log = o1.open('https://github.com').read().decode()
# POST
a_auth = re.findall('data-csrf="true" value="([^"]*)"', s_log)
m_data = {}
m_data['authenticity_token'] = a_auth[0]
m_data['value'] = 'aa540'
o2 = request.Request('https://github.com/signup_check/username')
o2.data = parse.urlencode(m_data).encode()
o2.headers['Accept'] = '*/*'
n = o1.open(o2).getcode()
print(n)
