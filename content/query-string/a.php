<?php
# example A
$s = 'one=odd&two=even';
parse_str($s, $mA);
# example B
$m = ['one' => 'odd', 'two' => 'even'];
$sB = http_build_query($m);
# print
var_dump($mA, $sB);
