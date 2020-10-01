<?php
# example 1
$m['year'] = 2019;
# example 2
$m2 = ['month' => 12];
$m = array_merge($m, $m2);
# example 3
$m += ['day' => 31];
# print
print_r($m);
