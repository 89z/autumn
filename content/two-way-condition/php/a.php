<?php
$n = 10;
# example 1
if ($n > 0) {
   $s1 = '+';
} else if ($n < 0) {
   $s1 = '-';
} else {
   $s1 = 'zero';
}
# example 2
if ($n > 0) {
   $s2 = '+';
} elseif ($n < 0) {
   $s2 = '-';
} else {
   $s2 = 'zero';
}
# example 3
$s3 = $n < 0 ? '-' : '+';
# print
var_dump($s1 == '+', $s2 == '+', $s3 == '+');
