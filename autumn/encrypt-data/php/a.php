<?php
extension_loaded('openssl') or die('openssl');

[$data, $key] = ['January February', '0123456789ABCDEF'];
$s = openssl_encrypt($data, 'aes-128-ecb', $key, OPENSSL_ZERO_PADDING);
var_dump($s == 'rc4qAleQ6vB5rbmug1qv5g==');
