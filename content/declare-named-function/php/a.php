<?php
# example 1
function f1(string $s, string $c): bool {
   return $s[0] == $c;
}
# example 2
function f2($s, $c) {
   return $s[0] == $c;
}
# print
$b1 = f1('June', 'J');
$b2 = f2('June', 'J');
var_dump($b1, $b2);
