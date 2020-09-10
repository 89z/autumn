<?php
# example 1
$s = 'one=odd&two=even';
parse_str($s, $m);
# example 2
$m2 = ['one' => 'odd', 'two' => 'even'];
$s2 = http_build_query($m2);
# print
var_dump($m, $s2);
