<?php
$m1 = ['Sun' => 10, 'Mon' => 11];
# example 1
$n1 = $m1['Mon'];
# example 2
$n2 = $m1{'Mon'};
# print
var_dump($n1, $n2);
