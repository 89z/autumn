<?php
$s1 = '/sunday/monday.tar.xz';
# example 1
$b1 = dirname($s1) == '/sunday';
# example 2
$b2 = pathinfo($s1, PATHINFO_DIRNAME) == '/sunday';
# print
var_dump($b1, $b2);
