<?php
$s = 'Sunday Monday';

# example 1
$n = preg_match('/..n/', 'Sunday Monday');
if ($n == 1) {
   echo $s, "\n";
}

# example 2
$n = preg_match('/..n/', 'Sunday Monday', $a);
if ($n == 1) {
   echo $a[0], "\n";
}
