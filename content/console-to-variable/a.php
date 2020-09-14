<?php
# example 1
$r = fopen('php://stdin', 'r');
$s = fgets($r);
# example 2
$s2 = fgets(STDIN);
# example 3
fscanf(STDIN, '%d', $n);
# print
var_dump($s, $s2, $n);
