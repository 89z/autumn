<?php
$a = [10, 11];
# example 1
array_push($a, 12);
# example 2
$a[] = 13;
# example 3
$a2 = [14, 15];
$a = array_merge($a, $a2);
# print
var_dump($a);
