s = '0a0b0c0d'
t = bytes.fromhex(s)
print(t == b'\n\v\f\r')
