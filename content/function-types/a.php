<?php
# example 1
function f(string $s): int {
   return strlen($s);
}
# example 2
$f2 = function (string $s): int {
   return strlen($s);
};
# example 3
$f3 = fn (string $s): int => strlen($s);
# print
$n = f('May');
$n2 = $f2('May');
$n3 = $f3('May');
var_dump($n == 3, $n2 == 3, $n3 == 3);
