<?php
# example 1
$s1 = file_get_contents('a.txt');
# example 2
$r1 = fopen('a.txt', 'r');
$s2 = stream_get_contents($r1);
# example 3
$a1 = file('a.txt');
# print
var_dump($s1, $s2, $a1);
