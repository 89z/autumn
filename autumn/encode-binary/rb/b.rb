require 'base64'
s = "\xa\xb\xc\xd"
t = Base64.strict_encode64(s)
puts t == 'CgsMDQ=='
