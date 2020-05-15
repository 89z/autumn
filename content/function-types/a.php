<?php
# example 1
function f1($s1) {
   return intval($s1);
}
$n1 = f1('10');
# example 2
$f2 = function ($s1) {
   return intval($s1);
};
$n2 = $f2('10');
# example 3
$f3 = fn($s1) => intval($s1);
$n3 = $f3('10');
# print
var_dump($n1, $n2, $n3);
