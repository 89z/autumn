<?php
$s1 = 'a.txt';
# example 1
$b1 = is_dir($s1);
# example 2
$b2 = file_exists($s1);
# example 3
$b3 = is_file($s1);
# print
var_dump($b1, $b2, $b3);
