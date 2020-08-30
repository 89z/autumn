<?php
$s = 'index.md';
# example 1
$s1 = realpath($s);
# example 2
$o = new SplFileInfo($s);
$s2 = $o->getRealPath();
# print
var_dump($s1, $s2);
