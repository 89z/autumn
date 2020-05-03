<?php
$s1 = 'https://example.com/one?two=even';
$m1 = parse_url($s1);
# example 1
$s2 = $m1['host'];
# example 2
$s3 = $m1['path'];
# example 3
$s4 = $m1['query'];
# print
var_dump($s2 == 'example.com', $s3 == '/one', $s4 == 'two=even');
