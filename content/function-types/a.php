<?php
# example 1
function f1($n1) {
   return $n1 + 10;
}
# example 2
$f2 = function ($n1) {
   return $n1 + 10;
};
# print
var_dump(f1(20), $f2(20));
