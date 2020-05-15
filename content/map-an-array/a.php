<?php
$a1 = ['Sun', 'Mon'];
# example 1
$a2 = array_map('strlen', $a1);
# example 2
$a3 = array_map(fn($s1) => strlen($s1), $a1);
# print
var_dump($a2, $a3);
