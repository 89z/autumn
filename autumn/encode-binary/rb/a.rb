require 'base64'
s = "\xa\xb\xc"
t = Base64.encode64(s)
puts t == "CgsM\n"
