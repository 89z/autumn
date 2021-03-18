s = "\xa\xb\xc\xd"
t = s.unpack('H*')[0]
puts t == '0a0b0c0d'
