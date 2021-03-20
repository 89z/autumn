require 'base64'
require 'openssl'
plain, key = 'January February', 'KDKDKDKDKDKDKDKD'

# example 1
cipher = OpenSSL::Cipher.new('aes-128-cbc').encrypt
cipher.iv = 'IVIVIVIVIVIVIVIV'
cipher.key = key
s = Base64.strict_encode64(cipher.update(plain) + cipher.final)

# example 2
cipher = OpenSSL::Cipher.new('aes-128-ecb').encrypt
cipher.key = key
t = Base64.strict_encode64(cipher.update(plain))

# print
puts(
   s == 'BvfnZp4jmCaveE6kefhumpZ0raWX9GDojfPasgSwLTM=',
   t == 'hr0e+xH2oi0mYHMmdGQCnQ=='
)
