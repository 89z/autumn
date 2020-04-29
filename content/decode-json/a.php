<?php
extension_loaded('json') or die('php-json');
$s1 = file_get_contents('a.json');
# example 1
$o1 = json_decode($s1);
$n1 = $o1->Sunday;
# example 2
$m1 = json_decode($s1, true);
$n2 = $m1['Sunday'];
# print
var_dump($n1, $n2);
