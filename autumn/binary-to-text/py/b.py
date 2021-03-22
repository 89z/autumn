import base64
s = b'\n\v\f\r'
t = base64.b64encode(s)
print(t == b'CgsMDQ==')
