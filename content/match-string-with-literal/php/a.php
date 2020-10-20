<?php
# example 1
function hasprefix($s, $s1) {
   return strpos($s, $s1) === 0;
}
# example 2
function contains($s, $s2) {
   return strpos($s, $s2) !== false;
}
# print
$b1 = hasprefix('June', 'Ju');
$b2 = contains('June', 'un');
var_dump($b1, $b2);
