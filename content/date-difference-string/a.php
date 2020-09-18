<?php
# example 1
$o = new DateTime('2019-12-31');
$o1 = new DateTime('2019-12-31T23:59:59');
$s1 = $o1->diff($o)->format('%H:%I:%S');
# example 2
$o = date_create('2019-12-31');
$o2 = date_create('2019-12-31T23:59:59');
$o2a = date_diff($o2, $o);
$s2 = date_interval_format($o2a, '%H:%I:%S');
# print
var_dump($s1 == '23:59:59', $s2 == '23:59:59');
