<?php
$s1 = '/sunday/monday.tar.xz';
# example 1
$b1 = dirname($s1) == '/sunday';
# example 2
$b2 = basename($s1) == 'monday.tar.xz';
# example 3
$b3 = pathinfo($s1, PATHINFO_DIRNAME) == '/sunday';
# example 4
$b4 = pathinfo($s1, PATHINFO_BASENAME) == 'monday.tar.xz';
# example 5
$b5 = pathinfo($s1, PATHINFO_FILENAME) == 'monday.tar';
# example 6
$b6 = pathinfo($s1, PATHINFO_EXTENSION) == 'xz';
# print
var_dump($b1, $b2, $b3, $b4, $b5, $b6);
