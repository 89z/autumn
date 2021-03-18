import base64
s = open('foo.png').read()
b64 = base64.b64encode(s)
s2 = base64.b64decode(b64)
