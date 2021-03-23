s = "\n\v\f\r"
t = s.unpack('H*')[0]
puts t == '0a0b0c0d'
