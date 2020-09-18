<?php
$m = parse_url('https://example.com/one?two=even');
# example A
$sA = $m['host'];
# example B
$sB = $m['path'];
# example C
$sC = $m['query'];
# print
var_dump($sA == 'example.com', $sB == '/one', $sC == 'two=even');
