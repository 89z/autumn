<?php
# example 1
function u1() {
   $a1 = func_get_args();
   return implode(' & ', $a1);
}
# example 2
function u2(...$a1) {
   return implode(' & ', $a1);
}
# print
$s1 = u1('sun', 'mon');
$s2 = u2('sun', 'mon');
var_dump($s1, $s2);
