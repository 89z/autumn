<?php
$n1 = 0;
# example 1
if ($n1 > 0) {
   $s1 = 'positive';
} else if ($n1 < 0) {
   $s1 = 'negative';
} else {
   $s1 = 'zero';
}
# example 2
if ($n1 > 0) {
   $s2 = 'positive';
} elseif ($n1 < 0) {
   $s2 = 'negative';
} else {
   $s2 = 'zero';
}
# example 3
$s3 = $n1 < 0 ? 'negative' : 'positive';
# print
var_dump($s1, $s2, $s3);
