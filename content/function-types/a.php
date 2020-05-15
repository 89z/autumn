<?php
# example 1
function f1($s1) {
   return strlen($s1);
}
$n1 = f1('Sunday');
# example 2
$f2 = function ($s1) {
   return strlen($s1);
};
$n2 = f1('Sunday');
# example 3
$f3 = fn($s1) => strlen($s1);
$n3 = $f3('Sunday');
# print
var_dump($n1, $n2, $n3);
