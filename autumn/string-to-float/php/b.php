<?php
$s = '9.9';
# example 1
$n1 = sscanf($s, '%f')[0];
# example 2
sscanf($s, '%f', $n2);
# print
var_dump($n1 === 9.9, $n2 === 9.9);
