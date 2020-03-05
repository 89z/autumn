<?php
# example 1
$s1 = 'one=odd&two=even';
parse_str($s1, $m1);
# example 2
$m2 = ['one' => 'odd', 'two' => 'even'];
$s2 = http_build_query($m2);
# print
var_dump($m1, $s2);
