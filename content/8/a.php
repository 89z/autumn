<?php
$s = '2019-12-31';
# example 1
$n = strtotime($s);
# example 2
$o = new DateTime($s);
$n2 = $o->getTimestamp();
# example 3
$o2 = DateTime::createFromFormat('!Y-m-d', $s);
$n3 = $o2->getTimestamp();
# print
var_dump($n, $n2, $n3);
