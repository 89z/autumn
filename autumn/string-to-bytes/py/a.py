import base64
s = '0A0B0C0D'
b = base64.b16decode(s)
print(b == b'\n\v\f\r')
