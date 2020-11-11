<?php
$m = parse_url('https://example.com/one?two=even');
# example 1
$s1 = $m['host'];
# example 2
$s2 = $m['query'];
# print
var_dump($s1 == 'example.com', $s2 == 'two=even');
