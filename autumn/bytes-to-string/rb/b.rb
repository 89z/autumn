require 'base64'
s = "\n\v\f\r"
t = Base64.strict_encode64(s)
puts t == 'CgsMDQ=='
