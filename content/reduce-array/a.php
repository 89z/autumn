<?php
$a = ['May', 'June'];
$f = fn ($s_acc, $s_cur) => $s_acc . $s_cur;
# example 1
$s = array_reduce($a, $f);
# example 2
$s2 = '';
foreach ($a as $s_cur) {
   $s2 = $f($s2, $s_cur);
}
# print
var_dump($s == 'MayJune', $s2 == 'MayJune');
