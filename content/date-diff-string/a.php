<?php
# example 1
$o1 = new DateTime('2019-12-31T00:00:00');
$o2 = new DateTime('2019-12-31T23:59:59');
$o3 = $o1->diff($o2);
$s1 = $o3->format('%H:%I:%S');
# example 2
$o1 = date_create('2019-12-31T00:00:00');
$o2 = date_create('2019-12-31T23:59:59');
$o3 = date_diff($o2, $o1);
$s2 = date_interval_format($o3, '%H:%I:%S');
# print
var_dump($s1, $s2);
