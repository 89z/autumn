<?php
# example 1
$a1 = ['Sun'];
array_push($a1, 'Mon');
# example 2
$a2 = ['Sun'];
$a2[] = 'Mon';
# example 3
$a3 = ['Sun'];
$a3 = array_merge($a3, ['Mon']);
# print
var_dump($a1, $a2, $a3);
