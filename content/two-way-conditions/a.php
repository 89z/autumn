<?php
$n1 = 10;
# example 1
$n2 = $n1 == 0 ? 20 : $n1;
# example 2
if ($n1 > 30) {
   $n3 = 30;
} elseif ($n1 > 20) {
   $n3 = 20;
} else {
   $n3 = $n1;
}
# print
var_dump($n2, $n3);
