import http.client, urllib.parse, urllib.request
def get_1(s_link):
   m1 = urllib.request
   return m1.urlopen(s_link).read().decode()
def get_2(s_link):
   m1 = urllib.parse
   m2 = http.client
   m3 = m1.urlparse(s_link)
   m4 = m2.HTTPConnection(m3.hostname)
   m4.request('GET', m3.path)
   return m4.getresponse().read().decode()
s1 = get_1('http://speedtest.lax.hivelocity.net')
s2 = get_2('http://speedtest.lax.hivelocity.net')
print(s1, s2)
