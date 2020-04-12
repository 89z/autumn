<?php
$s1 = '2020-03-29';
# example 1
$n1 = preg_match('/\d/', $s1);
# example 2
preg_match('/-../', $s1, $a1);
# example 3
preg_match_all('/-../', $s1, $a2);
# example 4
preg_match('/-(..)-(..)/', $s1, $a3);
# print
var_dump($n1, $a1, $a2, $a3);
