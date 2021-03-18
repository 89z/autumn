require 'base64'
require 'openssl'

def encrypt(data, key)
   cipher = OpenSSL::Cipher.new('aes-128-ecb').encrypt
   cipher.key = key
   return Base64.strict_encode64(cipher.update(data))
end

s = encrypt('January February', '0123456789ABCDEF')
puts s == 'rc4qAleQ6vB5rbmug1qv5g=='
