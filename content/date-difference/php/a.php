<?php
# example 1
$n = strtotime('2019-12-31');
$n1a = time();
$n1 = $n1a - $n;
# example 2
$o = new DateTime('2019-12-31');
$o2 = new DateTime;
$n2 = $o2->diff($o)->days;
# print
var_dump($n1, $n2);
