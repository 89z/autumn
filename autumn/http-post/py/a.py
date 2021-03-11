import httplib
import urllib

url = 'www.acme.com'
conn = httplib.HTTPConnection(url)
data = urllib.urlencode({
  'item': 'anvil',
  'qty': 1})
conn.request('POST', '/orders', data)
resp = conn.getresponse()
if resp.status == httplib.OK:
  s = resp.read()
