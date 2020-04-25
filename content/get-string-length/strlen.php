<?php
$s1 = '♠';
# example 1
$n1 = strlen($s1);
# example 2
$s2 = utf8_decode($s1);
$n2 = strlen($s2);
# print
var_dump($n1 == 3, $n2 == 1);
