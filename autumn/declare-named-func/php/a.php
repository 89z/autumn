<?php

# example 1
function f1(int $n, int $n1): bool {
   return $n > $n1;
}

# example 2
function f2($n, $n2) {
   return $n > $n2;
}

# print
$b1 = f1(9, 8);
$b2 = f2(9, 8);
var_dump($b1, $b2);
