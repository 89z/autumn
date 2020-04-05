<?php
# example 1
$o1 = simplexml_load_file('a.xml');
# example 2
$s1 = file_get_contents('a.xml');
$o2 = simplexml_load_string($s1);
# example 3
$o3 = $o1->Monday;
# example 4
$o4 = $o1['ten'];
# example 5
$o5 = $o1['ten'];
$s2 = (string)($o5);
# print
var_dump($o1, $o2, $o3, $o4, $s2);
