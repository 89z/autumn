<?php

# example 1
$n = strtotime('2019-12-31');
$n2 = time();
$n3 = $n2 - $n;

# example 2
$o = new DateTime('2019-12-31');
$o2 = new DateTime;
$n4 = $o2->diff($o)->days;

# print
var_dump($n3, $n4);
