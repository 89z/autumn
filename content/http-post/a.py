from http import cookiejar
from urllib import request, parse
import re
o1 = request.build_opener(
   request.HTTPSHandler(debuglevel=1),
   request.HTTPCookieProcessor(cookiejar.CookieJar())
)
# GET
s_log = o1.open('https://github.com').read().decode()
# POST
a_auth = re.findall('data-csrf="true" value="([^"]*)"', s_log)
m_data = {}
m_data['value'] = 'aa540'
m_data['authenticity_token'] = a_auth[0]
s1 = parse.urlencode(m_data)
s2 = s1.encode()
o2 = request.Request(
   'https://github.com/signup_check/username',
   data=s2,
   headers={'Accept': '*/*'}
)
o1.open(o2)
