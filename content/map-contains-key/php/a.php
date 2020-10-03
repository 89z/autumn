<?php
$m['year'] = 2019;
# example 1
$b1 = array_key_exists('year', $m);
# example 2
$b2 = key_exists('year', $m);
# print
var_dump($b1, $b2);
