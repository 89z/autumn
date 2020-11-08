<?php
$s = 'https://example.com/one?two=even';
# example 1
$s1 = parse_url($s, PHP_URL_HOST);
# example 2
$s2 = parse_url($s, PHP_URL_PATH);
# example 3
$s3 = parse_url($s, PHP_URL_QUERY);
# print
var_dump($s1 == 'example.com', $s2 == '/one', $s3 == 'two=even');
