<?php
# example 1
$n1 = time();
# example 2
$o1 = date_create('2020-05-23');
$n2 = date_timestamp_get($o1);
# example 3
$n3 = strtotime('2020-05-23');
# print
var_dump($n1, $n2, $n3);
