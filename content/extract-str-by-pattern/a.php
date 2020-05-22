<?php
$s1 = 'Wednesday';
# example 1
preg_match('/e./', $s1, $a1);
# example 2
preg_match_all('/e./', $s1, $a2);
# example 3
preg_match('/e(..)/', $s1, $a3);
# example 4
preg_match_all('/e(..)/', $s1, $a4);
# print
var_dump($a1, $a2, $a3, $a4);
