<?php

# example 1
$f1 = function (int $n, int $n2): bool {
   return $n > $n2;
};

# example 2
$f2 = function ($n, $n2) {
   return $n > $n2;
};

# example 3
$f3 = fn (int $n, int $n2): bool => $n > $n2;

# example 4
$f4 = fn ($n, $n2) => $n > $n2;

# print
$b1 = $f1(9, 8);
$b2 = $f2(9, 8);
$b3 = $f3(9, 8);
$b4 = $f4(9, 8);
var_dump($b1, $b2, $b3, $b4);
