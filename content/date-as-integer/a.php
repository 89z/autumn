<?php
# example 1
$n1 = time();
# example 2
$n2 = strtotime('2020-05-23');
# example 3
$o1 = date_create('2020-05-23');
$n3 = date_timestamp_get($o1);
# print
var_dump($n1, $n2, $n3);
