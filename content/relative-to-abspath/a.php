<?php
$s = 'index.md';
# example 1
$s2 = realpath($s);
# example 2
$o = new SplFileInfo($s);
$s3 = $o->getRealPath();
# print
var_dump($s2, $s3);
