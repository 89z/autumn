<?php
$s = 'index.md';
# example 1
$b1 = file_exists($s);
# example 2
$b2 = is_file($s);
# example 3
$b3 = is_dir($s);
# print
var_dump($b1, $b2, $b3);
