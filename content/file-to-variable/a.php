<?php
# example 1
$s1 = file_get_contents('a.txt');
# example 2
$a1 = file('a.txt');
# example 3
$a2 = file('a.txt', FILE_IGNORE_NEW_LINES);
# example 4
$a3 = file('a.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);
# example 5
$r1 = fopen('a.txt', 'r');
$s2 = stream_get_contents($r1);
# print
var_dump($s1, $a1, $a2, $a3, $r2);
