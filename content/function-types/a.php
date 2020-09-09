<?php
# example 1
function f1(string $s): int {
   return strlen($s);
}
# example 2
$f2 = function (string $s): int {
   return strlen($s);
};
# example 3
$f3 = fn (string $s): int => strlen($s);
# print
$n1 = f1('Sunday');
$n2 = $f2('Sunday');
$n3 = $f3('Sunday');
var_dump($n1, $n2, $n3);
