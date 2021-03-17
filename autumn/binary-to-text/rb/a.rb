require 'base64'
s = File.open("foo.png").read
b64 = Base64.encode64(s)
s2 = Base64.decode64(b64)
