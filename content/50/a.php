<?php
# example 1
$a1 = [10, 11];
array_push($a1, 12);
# example 2
$a2 = [10, 11];
$a2[] = 12;
# example 3
$a3 = [10, 11];
$a3 = array_merge($a3, [12, 13]);
# print
var_dump($a1, $a2, $a3);
