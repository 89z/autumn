import base64
b = b'\n\v\f\r'
c = base64.b16encode(b)
print(c == b'0A0B0C0D')
