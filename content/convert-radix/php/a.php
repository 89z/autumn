<?php
# example 1
$n = 1577858399;
$s1 = base_convert($n, 10, 36);
# example 2
$s = 'q3ezbz';
$n2 = +base_convert($s, 36, 10);
# print
var_dump($s1 === $s, $n2 === $n);
