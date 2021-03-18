require 'base64'
require 'openssl'
s = 'January February'
c = OpenSSL::Cipher.new('aes-128-ecb').encrypt
c.key = '0123456789ABCDEF'
t = c.update(s)
puts Base64.strict_encode64(t) == 'rc4qAleQ6vB5rbmug1qv5g=='
