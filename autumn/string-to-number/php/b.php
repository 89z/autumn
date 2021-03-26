<?php
$s = '100,200';
# example 1
$a = sscanf($s, '%d,%d');
# example 2
sscanf($s, '%d,%d', $d, $e);
# print
var_dump($a, $d, $e);
