<?php
$s = '10,11';
# example 1
$a = sscanf($s, '%d,%d');
# example 2
sscanf($s, '%d,%d', $n, $n2);
# print
var_dump($a, $n, $n2);
