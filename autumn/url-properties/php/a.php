<?php
$s = 'http://php.net/function.parse-url?west=left';

# example 1
$t = parse_url($s, PHP_URL_HOST);
var_dump($t == 'php.net');

# example 2
$t = parse_url($s, PHP_URL_PATH);
var_dump($t == '/function.parse-url');

# example 3
$t = parse_url($s, PHP_URL_QUERY);
var_dump($t == 'west=left');
