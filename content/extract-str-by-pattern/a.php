<?php
$s1 = 'package';
# example 1
preg_match('/p./', $s1, $a1);
# example 2
preg_match_all('/a./', $s1, $a2);
# example 3
preg_match('/p(..)/', $s1, $a3);
# print
var_dump($a1, $a2, $a3);
