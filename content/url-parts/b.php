<?php
$m = parse_url('https://example.com/one?two=even');
# example 1
$s = $m['host'];
# example 2
$s2 = $m['path'];
# example 3
$s3 = $m['query'];
# print
var_dump($s == 'example.com', $s2 == '/one', $s3 == 'two=even');
