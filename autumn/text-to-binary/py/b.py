import base64
s = 'CgsMDQ=='
t = base64.b64decode(s)
print(t == b'\n\v\f\r')
