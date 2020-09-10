<?php
$s = 'https://example.com/one?two=even';
# example 1
$s2 = parse_url($s, PHP_URL_HOST);
# example 2
$s3 = parse_url($s, PHP_URL_PATH);
# example 3
$s4 = parse_url($s, PHP_URL_QUERY);
# print
var_dump($s2 == 'example.com', $s3 == '/one', $s4 == 'two=even');
