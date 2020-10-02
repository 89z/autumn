<?php
# example 1
$s = 'one=odd&two=even';
parse_str($s, $m1);
# example 2
$m = ['one' => 'odd', 'two' => 'even'];
$s2 = http_build_query($m);
# print
var_dump($m1, $s2);
