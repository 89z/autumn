<?php
$s = 'March';
# example 1
$s1 = substr($s, 2, 1);
# example 2
$s2 = substr($s, 2);
# print
var_dump($s1 == 'r', $s2 == 'rch');
