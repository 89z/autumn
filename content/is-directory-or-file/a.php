<?php
$s = 'index.md';
# example 1
$b = is_file($s);
# example 2
$b2 = ! is_dir($s);
# print
var_dump($b, $b2);
