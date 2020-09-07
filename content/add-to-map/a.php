<?php
# example 1
$m['year'] = 2020;
# example 2
$m2 = ['month' => 9];
$m = array_merge($m, $m2);
# example 3
$m += ['day' => 7];
# print
print_r($m);
