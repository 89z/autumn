require 'base64'
s = 'IyQl'
t = Base64.decode64(s)
puts t == '#$%'
