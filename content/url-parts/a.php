<?php
$s1 = 'http://sun.com/mon?tue=10';
# example 1
$m1 = parse_url($s1);
# example 2
$s2 = parse_url($s1, PHP_URL_QUERY);
# print
var_dump($m1, $s2);
