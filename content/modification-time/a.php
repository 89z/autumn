<?php
$s = 'index.md';
# example 1
$n1 = time();
touch($s, $n1);
# example 2
$n2 = filemtime($s);
# print
var_dump($n1 == $n2);
