<?php
# example 1
function f1() {
   $a1 = func_get_args();
   return implode(',', $a1);
}
$s1 = f1('Sun', 'Mon');
# example 2
function f2(...$a1) {
   return implode(',', $a1);
}
$s2 = f2('Sun', 'Mon');
# print
var_dump($s1, $s2);
