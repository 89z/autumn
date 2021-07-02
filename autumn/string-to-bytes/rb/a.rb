s = '0a0b0c0d'
t = [s].pack('H*')
puts t == "\n\v\f\r"
