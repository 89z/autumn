<?php
$s = 'https://example.com/one?two=even';
# example A
$sA = parse_url($s, PHP_URL_HOST);
# example B
$sB = parse_url($s, PHP_URL_PATH);
# example C
$sC = parse_url($s, PHP_URL_QUERY);
# print
var_dump($sA == 'example.com', $sB == '/one', $sC == 'two=even');
