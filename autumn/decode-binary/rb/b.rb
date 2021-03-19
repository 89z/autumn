require 'base64'
s = 'CgsMDQ=='
t = Base64.decode64(s)
puts t == "\n\v\f\r"
