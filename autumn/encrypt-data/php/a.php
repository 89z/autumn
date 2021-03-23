<?php
[$plain, $key] = ['January February', 'KDKDKDKDKDKDKDKD'];

# example 1
$s = openssl_encrypt($plain, 'aes-128-cbc', $key, iv: 'IVIVIVIVIVIVIVIV');
var_dump($s == 'BvfnZp4jmCaveE6kefhumpZ0raWX9GDojfPasgSwLTM=');

# example 2
$s = openssl_encrypt($plain, 'aes-128-ecb', $key, OPENSSL_ZERO_PADDING);
var_dump($s == 'hr0e+xH2oi0mYHMmdGQCnQ==');

# example 3
$s = openssl_encrypt($plain, 'bf-cbc', $key, iv: 'IVIVIVIV');
var_dump($s == 'zH/Z28yXchZ8wih5jErxTgKJpPxLcHgo');
