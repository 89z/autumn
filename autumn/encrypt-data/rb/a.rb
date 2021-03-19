require 'base64'
require 'openssl'
data, key = 'January February', 'PassPassPassPass'

# example 1
cipher = OpenSSL::Cipher.new('aes-128-cbc').encrypt
cipher.iv = 'IvIvIvIvIvIvIvIv'
cipher.key = key
s1 = Base64.strict_encode64(cipher.update(data) + cipher.final)

# example 2
cipher = OpenSSL::Cipher.new('aes-128-ecb').encrypt
cipher.key = key
s2 = Base64.strict_encode64(cipher.update(data))

# print
puts(
   s1 == 'x2B4PSY4exW+U7x/jmOkSJmDsB0IgpsLeHSICOQnPF8=',
   s2 == 'LxPTC0HIdnxnCj54rzkEjA=='
)
