<?php
$n = 10;
# example 1
if ($n > 12) {
   $s1 = 'Tue';
} else if ($n > 11) {
   $s1 = 'Mon';
} else {
   $s1 = 'Sun';
}
# example 2
if ($n > 12) {
   $s2 = 'Tue';
} elseif ($n > 11) {
   $s2 = 'Mon';
} else {
   $s2 = 'Sun';
}
# example 3
$s3 = $n > 11 ? 'Mon' : 'Sun';
# print
var_dump($s1, $s2, $s3);
