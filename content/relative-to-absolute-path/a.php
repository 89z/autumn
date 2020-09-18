<?php
$s = 'index.md';
# example A
$sA = realpath($s);
# example B
$o = new SplFileInfo($s);
$sB = $o->getRealPath();
# print
var_dump($sA, $sB);
