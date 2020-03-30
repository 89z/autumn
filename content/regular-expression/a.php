<?php
$s1 = '2020-03-29';
# example 1
preg_match('/-../', $s1, $a1);
# example 2
preg_match_all('/-../', $s1, $a2);
# example 3
preg_match('/-(..)-(..)/', $s1, $a3);
# print
var_dump($a1, $a2, $a3);
