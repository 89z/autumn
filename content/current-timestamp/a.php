<?php
# example 1
$n1 = time();
# example 2
$n2 = strtotime('now');
# example 3
$o1 = date_create('now');
$n3 = date_timestamp_get($o1);
# print
var_dump($n1, $n2, $n3);
