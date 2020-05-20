<?php
# example 1
$r1 = fopen('php://stdin', 'r');
$s1 = fgets($r1);
# example 2
$s2 = fgets(STDIN);
# example 3
fscanf(STDIN, '%d', $n1);
# print
var_dump($s1, $s2, $n1);
