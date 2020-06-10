<?php
$s = '2019-12-31';
# example 1
$o1 = new DateTime($s);
# example 2
$o2 = date_create($s);
# print
var_dump($o1, $o2);
