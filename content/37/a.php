<?php
$n = 1;
# example 1
if ($n > 0) {
   $s = '+';
} else if ($n < 0) {
   $s = '-';
} else {
   $s = 'zero';
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
var_dump($s == '+', $s2 == '+', $s3 == '+');
