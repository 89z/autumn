<?php
extension_loaded('openssl') or die('openssl');
[$data, $pass] = ['January February', 'PassPassPassPass'];

# example 1
$s1 = openssl_encrypt($data, 'aes-128-cbc', $pass, iv: 'IvIvIvIvIvIvIvIv');

# exmaple 2
$s2 = openssl_encrypt($data, 'aes-128-ecb', $pass, OPENSSL_ZERO_PADDING);

# print
var_dump(
   $s1 == 'x2B4PSY4exW+U7x/jmOkSJmDsB0IgpsLeHSICOQnPF8=',
   $s2 == 'LxPTC0HIdnxnCj54rzkEjA=='
);
