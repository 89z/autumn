require 'openssl'

pp OpenSSL::Cipher.ciphers
=begin
[
   "aes-128-ecb",
   "aes-192-ecb",
   "aes-256-ecb",
]
=end
