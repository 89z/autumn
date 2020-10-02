<?php
$s = 'php -v';
# example 1
$s1 = `$s`;
# example 2
$s2 = shell_exec($s);
# example 3
$r = popen($s, 'r');
$s3 = stream_get_contents($r);
# example 4
exec($s, $a4);
# print
var_dump($s1, $s2, $s3, $a4);
