<?php
[$plain, $key] = ['north east south', 'KDKDKDKDKDKDKDKD'];

# example 1
$s = openssl_encrypt($plain, 'aes-128-ecb', $key, OPENSSL_ZERO_PADDING);
var_dump($s == '4HbK4aRsiiLXEtT99V2Xgg==');

# example 2
$s = openssl_encrypt($plain, 'bf-cbc', $key, iv: 'IVIVIVIV');
var_dump($s == 'KQCQKBh5gosLsrIXqvB5tASOMDv4jVpM');
