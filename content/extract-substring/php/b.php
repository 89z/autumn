<?php
$s = 'May';
# example 1
$s1 = substr($s, 1, 1);
# example 2
$s2 = substr($s, 1);
# print
var_dump($s1 == 'a', $s2 == 'ay');
