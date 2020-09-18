<?php
$a = ['May', 'June'];
# example A
$sA = array_reduce($a, fn ($sc, $sd) => $sc . $sd);
# example B
$f = fn ($sc, $sd) => $sc . $sd;
$sB = array_reduce($a, $f);
# print
var_dump($sA == 'MayJune', $sB == 'MayJune');
