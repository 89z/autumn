<?php
# example 1
$a1 = ['Sun', 'Mon'];
sort($a1);
# example 2
$a2 = ['Mon', 'Tue'];
rsort($a2);
# example 3
$a3 = ['Monday', 'Tue'];
array_multisort(array_map('strlen', $a3), $a3);
# example 4
$a4 = ['Monday', 'Tue'];
usort($a4, fn($s1, $s2) => strlen($s1) - strlen($s2));
# print
var_dump($a1, $a2, $a3, $a4);
