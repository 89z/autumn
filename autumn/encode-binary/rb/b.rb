require 'base64'
s = "\xa\xb\xc"
t = Base64.strict_encode64(s)
puts t == 'CgsM'
