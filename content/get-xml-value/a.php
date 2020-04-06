<?php
$o1 = simplexml_load_file('a.xml');
# example 1
$o2 = $o1->Monday;
# example 2
$o3 = $o1['ten'];
# example 3
$o4 = $o1['ten'];
$s1 = (string)($o4);
# print
var_dump($o2, $o3, $s1);
