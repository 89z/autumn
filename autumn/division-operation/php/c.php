<?php
[$f, $i] = [7.5, 7];
# example 1
$n1 = $f / 2;
# example 2
$n2 = (int)($f / 2);
# example 3
$n3 = (int)($i / 2);
# example 4
$n4 = $i / 2;
# print
var_dump($n1 == 3.75, $n2 == 3, $n3 == 3, $n4 == 3.5);
