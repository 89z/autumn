<?php
$s = 'Wednesday';
# example 1
preg_match('/e./', $s, $a);
# example 2
preg_match_all('/e./', $s, $a2);
# example 3
preg_match('/e(..)/', $s, $a3);
# example 4
preg_match_all('/e(..)/', $s, $a4);
# print
var_dump($a, $a2, $a3, $a4);
