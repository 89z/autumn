import base64
s = b'\n\v\f\r'
t = s.hex()
print(t == '0a0b0c0d')
