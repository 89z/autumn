<?php
$n = 10;
# example A
if ($n > 0) {
   $sA = '+';
} else if ($n < 0) {
   $sA = '-';
} else {
   $sA = 'zero';
}
# example B
if ($n > 0) {
   $sB = '+';
} elseif ($n < 0) {
   $sB = '-';
} else {
   $sB = 'zero';
}
# example C
$sC = $n < 0 ? '-' : '+';
# print
var_dump($sA == '+', $sB == '+', $sC == '+');
