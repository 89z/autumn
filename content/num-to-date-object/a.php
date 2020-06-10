<?php
$n = 365 * 24 * 60 * 60;
# example 1
$o1 = new DateTime;
$o1->setTimestamp($n);
# example 2
$o2 = date_create();
date_timestamp_set($o2, $n);
# print
var_dump($o1, $o2);
