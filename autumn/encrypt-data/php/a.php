<?php
extension_loaded('openssl') or die('openssl');
[$plain, $key] = ['January February', 'KDKDKDKDKDKDKDKD'];

# example 1
$s = openssl_encrypt($plain, 'aes-128-cbc', $key, iv: 'IVIVIVIVIVIVIVIV');

# exmaple 2
$t = openssl_encrypt($plain, 'aes-128-ecb', $key, OPENSSL_ZERO_PADDING);

# print
var_dump(
   $s == 'BvfnZp4jmCaveE6kefhumpZ0raWX9GDojfPasgSwLTM=',
   $t == 'hr0e+xH2oi0mYHMmdGQCnQ=='
);
