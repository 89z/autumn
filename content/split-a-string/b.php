<?php
$s = '8,9';
# example A
$a = sscanf($s, '%d,%d');
# example B
sscanf($s, '%d,%d', $nA, $nB);
# print
var_dump($a, $nA, $nB);
