require 'base64'
s = "\n\v\f\r"
t = Base64.encode64(s)
puts t == "CgsMDQ==\n"
