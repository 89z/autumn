<?php
$s = 'Wednesday';
# example 1
$n1 = preg_match('/^W/', $s);
# example 2
$n2 = preg_match('/we/i', $s);
# example 3
$n3 = preg_match('/(?i)we/', $s);
# example 4
$n4 = preg_match_all('/e./', $s);
# print
var_dump($n1 !== 0, $n2 !== 0, $n3 !== 0, $n4 !== 0);
