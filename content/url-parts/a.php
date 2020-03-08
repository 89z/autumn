<?php
$s1 = 'https://example.com/sun?mon=10';
# example 1
$m1 = parse_url($s1);
# example 2
$s2 = parse_url($s1, PHP_URL_PATH);
# example 3
$s3 = parse_url($s1, PHP_URL_QUERY);
# example 4
$s4 = parse_url($s1)['path'];
# example 5
$s5 = parse_url($s1)['query'];
# print
var_dump($m1, $s2, $s3, $s4, $s5);
