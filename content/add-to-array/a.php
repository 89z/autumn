<?php
# example 1
$a1 = [10, 20];
array_push($a1, 30);
# example 2
$a2 = [10, 20];
$a2[] = 30;
# example 3
$a3 = [10, 20];
$a3 = array_merge($a3, [30, 40]);
# print
var_dump($a1, $a2, $a3);
