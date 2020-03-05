<?php
$s1 = '{"sunday": 10, "monday": 11}';
# example 1
$o1 = json_decode($s1);
$n1 = $o1->sunday;
# example 2
$m1 = json_decode($s1, true);
$n2 = $m1['sunday'];
# print
var_dump($n1, $n2);
