<?php
# example 1
$r = fopen('php://stdin', 'r');
$s1 = fgets($r);
# example 2
$s2 = fgets(STDIN);
# example 3
fscanf(STDIN, '%d', $n3);
# print
var_dump($s1, $s2, $n3);
