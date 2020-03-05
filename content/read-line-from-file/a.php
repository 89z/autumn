<?php
$r1 = fopen('a.txt', 'r');
# example 1
$s1 = fgets($r1);
# example 2
$s2 = stream_get_line($r1, 0, "\n");
# print
var_dump($s1, $s2);
