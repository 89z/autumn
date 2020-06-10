<?php
$s = '2019-12-31';
# example 1
$n1 = strtotime($s);
# example 2
$o = new DateTime;
$n2 = $o->getTimestamp();
# print
var_dump($n1, $n2);
