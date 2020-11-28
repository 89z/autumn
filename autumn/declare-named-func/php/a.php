<?php
# example 1
function add(int $n, int $n1): int {
   return $n + $n1;
}

# example 2
function sub($n, $n2) {
   return $n - $n2;
}

# print
$n1 = add(8, 1);
$n2 = sub(8, 1);
var_dump($n1 == 9, $n2 == 7);
