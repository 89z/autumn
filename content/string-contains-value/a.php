<?php
# example 1
$s1 = 'Sunday'[0];
var_dump($s1 == 'S');
# example 2
$n1 = strpos('Sunday', 'Sun');
var_dump($n1 === 0);
# example 3
$n2 = strpos('Sunday', 'day');
var_dump($n2 !== false);
