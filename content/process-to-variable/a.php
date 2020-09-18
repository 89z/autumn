<?php
$s = 'php -v';
# example A
$sA = `$s`;
# example B
$sB = shell_exec($s);
# example C
$r = popen($s, 'r');
$sC = stream_get_contents($r);
# example D
exec($s, $aD);
# print
var_dump($sA, $sB, $sC, $aD);
