<?php
# example 1
function f1($s1) {
   return strlen($s1);
}
# example 2
$f2 = function ($s1) {
   return strlen($s1);
};
# example 3
$f3 = fn($s1) => strlen($s1);
# print
$n1 = f1('Sunday');
$n2 = $f2('Sunday');
$n3 = $f3('Sunday');
var_dump($n1, $n2, $n3);
