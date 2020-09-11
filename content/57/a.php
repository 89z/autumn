<?php
# example 1
$n = time();
# example 2
$n2 = strtotime('now');
# example 3
$o = new DateTime('now');
$n3 = $o->getTimestamp();
# example 4
$o = date_create('now');
$n4 = date_timestamp_get($o);
# print
var_dump($n, $n2, $n3, $n4);
