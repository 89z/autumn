<?php
$s1 = 'ag -V';
# example 1
$s2 = `$s1`;
# example 2
$s3 = shell_exec($s1);
# example 3
$r1 = popen($s1, 'r');
$s4 = stream_get_contents($r1);
# example 4
exec($s1, $a1);
# print
var_dump($s2, $s3, $s4, $a1);
