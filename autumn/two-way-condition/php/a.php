<?php
$n = 10;

# example 1
if ($n > 0) {
   $s = '+';
} else if ($n < 0) {
   $s = '-';
} else {
   $s = 'zero';
}
var_dump($s == '+');

# example 2
if ($n > 0) {
   $s = '+';
} elseif ($n < 0) {
   $s = '-';
} else {
   $s = 'zero';
}
var_dump($s == '+');
