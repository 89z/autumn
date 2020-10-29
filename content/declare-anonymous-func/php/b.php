<?php
# example 1
$f1 = fn (int $n, int $n1): bool => $n > $n1;
# example 2
$f2 = fn ($n, $n2) => $n > $n2;
# print
$b1 = $f1(9, 8);
$b2 = $f2(9, 8);
var_dump($b1, $b2);
