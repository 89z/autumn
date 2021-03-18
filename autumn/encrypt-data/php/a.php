<?php
$s = 'January February';

$t = openssl_encrypt(
   $s, 'aes-128-ecb', '0123456789ABCDEF', OPENSSL_ZERO_PADDING
);

var_dump($t == 'rc4qAleQ6vB5rbmug1qv5g==');
