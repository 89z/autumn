<?php
$a = ['May', 'June', 'July'];
# example A
$sA = $a[0];
# example B
$sB = reset($a);
# example C
$sC = end($a);
# example D
$aD = array_slice($a, 0, 2);
# example E
$aE = array_slice($a, -2);
# print
var_dump($sA, $sB, $sC, $aD, $aE);
