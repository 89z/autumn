<?php
$s = '100,200';
# example 1
$a1 = sscanf($s, '%d,%d');
# example 2
sscanf($s, '%d,%d', $n2, $n2a);
# print
var_dump($a1, $n2, $n2a);
