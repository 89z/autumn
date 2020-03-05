<?php
# example 1
$a1 = ['sun', 'mon'];
sort($a1);
# example 2
$a2 = ['mon', 'tue'];
rsort($a2);
# example 3
$a3 = ['monday', 'tue'];
array_multisort(array_map('strlen', $a3), $a3);
# example 4
$a4 = ['monday', 'tue'];
usort($a4, function ($s1, $s2) {
   return strlen($s1) - strlen($s2);
});
# print
var_dump($a1, $a2, $a3, $a4);
