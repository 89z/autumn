<?php
$s1 = 'Sunday';
# example 1
$n1 = preg_match('/^Su/', $s1);
# example 2
$n2 = preg_match('/un./', $s1);
# example 3
$n3 = preg_match('/ay$/', $s1);
# example 4
$n4 = preg_match('/su/i', $s1);
# example 5
$n5 = preg_match('/(?i)su/', $s1);
# print
var_dump($n1 == 1, $n2 == 1, $n3 == 1, $n4 == 1, $n5 == 1);
