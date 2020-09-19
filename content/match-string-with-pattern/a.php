<?php
$s = 'January';
# example 1
$n1 = preg_match('/^J/', $s);
# example 2
$n2 = preg_match('/ja/i', $s);
# example 3
$n3 = preg_match('/(?i)ja/', $s);
# example 4
$n3 = preg_match_all('/a./', $s);
# print
var_dump($n1 !== 0, $n2 !== 0, $n3 !== 0, $n4 !== 0);
