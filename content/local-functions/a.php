<?php
# example 1
function f(int $n, int $n2): bool {
   return $n > $n2;
}
# example 2
$f2 = function (int $n, int $n2): bool {
   return $n > $n2;
};
# example 3
$f3 = fn (int $n, int $n2): bool => $n > $n2;
# print
$b = f(9, 8);
$b2 = $f2(9, 8);
$b3 = $f3(9, 8);
var_dump($b, $b2, $b3);
