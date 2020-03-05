<?php
$m1 = ['sun' => 10, 'mon' => 11];
# example 1
$n1 = $m1['mon'];
# example 2
$n2 = $m1{'mon'};
# print
var_dump($n1, $n2);
