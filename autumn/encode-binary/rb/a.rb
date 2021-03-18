require 'base64'
s = "\xa\xb\xc\xd"
t = Base64.encode64(s)
puts t == "CgsMDQ==\n"
