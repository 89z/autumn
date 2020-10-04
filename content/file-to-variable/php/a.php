<?php
# example 1
$s1 = file_get_contents('a.txt');
# example 2
$r = fopen('a.txt', 'r');
$s2 = stream_get_contents($r);
# example 3
$a3 = file('a.txt');
# example 4
$a4 = file('a.txt', FILE_IGNORE_NEW_LINES);
# example 5
$a5 = file('a.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);
# print
var_dump($s1, $s2, $a3, $a4, $a5);
