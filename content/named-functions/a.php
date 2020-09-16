<?php
# example 1
function f($n, $n2) {
   return $n > $n2;
}
# example 2
function f2(int $n, int $n2): bool {
   return $n > $n2;
}
# print
$b = f(9, 8);
$b2 = f2(9, 8);
var_dump($b, $b2);
