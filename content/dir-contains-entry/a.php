<?php
$s1 = 'a.txt';
# example 1
$b1 = file_exists($s1);
# example 3
$b2 = is_file($s1);
# print
var_dump($b1, $b2);
