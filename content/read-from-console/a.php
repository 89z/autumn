<?php
# example 1
$r1 = fopen('php://stdin', 'r');
$s1 = fgets($r1);
# example 2
$s2 = fgets(STDIN);
# print
var_dump($s1, $s2);
