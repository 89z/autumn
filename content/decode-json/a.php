<?php
$s1 = '{"Sun": 10, "Mon": 11}';
# example 1
$o1 = json_decode($s1);
$n1 = $o1->Sun;
# example 2
$m1 = json_decode($s1, true);
$n2 = $m1['Sun'];
# print
var_dump($n1, $n2);
