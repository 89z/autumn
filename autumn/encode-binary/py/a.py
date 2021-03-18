import base64
s = b'\x0a\x0b\x0c\x0d'
t = base64.b64encode(s)
print(t == b'CgsMDQ==')
