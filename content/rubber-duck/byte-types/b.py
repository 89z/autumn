import http.client as o1
import urllib.parse as o2
s1 = 'http://speedtest.lax.hivelocity.net'
o3 = o2.urlparse(s1)
o4 = o1.HTTPConnection(o3.hostname)
o4.request('GET', o3.path)
o5 = o4.getresponse()
a1 = o5.read()
print(a1)