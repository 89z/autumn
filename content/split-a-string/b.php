<?php
$s1 = '10,11';
# example 1
$a1 = sscanf($s1, '%d,%d');
# example 2
sscanf($s1, '%d,%d', $n1, $n2);
# print
var_dump($a1, $n1, $n2);
