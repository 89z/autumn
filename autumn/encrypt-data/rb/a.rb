require 'base64'
require 'openssl'
plain, key = 'north east south', 'KDKDKDKDKDKDKDKD'

# example 1
cipher = OpenSSL::Cipher.new('aes-128-cbc').encrypt
cipher.iv = 'IVIVIVIVIVIVIVIV'
cipher.key = key
s = Base64.strict_encode64(cipher.update(plain) + cipher.final)
p s == 'gGPdBt1qCJCbHye4HX1E9ZUhYmjPsvEi7AsIs3XOReg='

# example 2
cipher = OpenSSL::Cipher.new('aes-128-ecb').encrypt
cipher.key = key
s = Base64.strict_encode64(cipher.update(plain))
p s == '4HbK4aRsiiLXEtT99V2Xgg=='
