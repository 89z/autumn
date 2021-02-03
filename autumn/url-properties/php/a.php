<?php
declare(strict_types = 1);
$s = 'http://php.net/function.parse-url?month=May';
# example 1
$s1 = parse_url($s, PHP_URL_HOST);
# example 2
$s2 = parse_url($s, PHP_URL_PATH);
# example 3
$s3 = parse_url($s, PHP_URL_QUERY);
# print
var_dump($s1, $s2, $s3);
