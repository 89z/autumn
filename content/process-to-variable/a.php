<?php
$s = 'php -v';
# example 1
$s2 = `$s`;
# example 2
$s3 = shell_exec($s);
# example 3
$r = popen($s, 'r');
$s4 = stream_get_contents($r);
# example 4
exec($s, $a);
# print
var_dump($s2, $s3, $s4, $a);
