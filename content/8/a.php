<?php
$s = '2019-12-31';
# example 1
$n1 = strtotime($s);
# example 2
$o1 = new DateTime($s);
$n2 = $o1->getTimestamp();
# example 3
$o2 = DateTime::createFromFormat('!Y-m-d', $s);
$n3 = $o2->getTimestamp();
# print
var_dump($n1, $n2, $n3);
