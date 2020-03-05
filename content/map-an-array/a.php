<?php
$a1 = [16, 25];
# example 1
$a2 = array_map('sqrt', $a1);
# example 2
$a3 = array_map(function ($n1) {
   return sqrt($n1);
}, $a1);
# print
print_r($a2);
print_r($a3);
