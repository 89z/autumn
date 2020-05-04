<?php
# example 1
$n1 = preg_match('/^Sun/', 'Sunday');
# example 2
$n2 = preg_match('/day/', 'Sunday');
# example 3
$n3 = preg_match('/day$/', 'Sunday');
# print
var_dump($n1, $n2, $n3);
