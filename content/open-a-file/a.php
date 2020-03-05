<?php
# read
$r1 = fopen('a.txt', 'r');
# read and write
$r2 = fopen('a.txt', 'r+');
# write
$r3 = fopen('a.txt', 'w');
# print
var_dump($r1, $r2, $r3);
