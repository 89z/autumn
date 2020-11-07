<?php
# example 1
$s1 = getenv('USERPROFILE');
# example 2
$m = getenv();
$s2 = $m['USERPROFILE'];
# print
$s = 'C:\Users\Steven';
var_dump($s1 == $s, $s2 == $s);
