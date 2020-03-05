<?php
# example 1
$s1 = `cal`;
# example 2
$s2 = shell_exec('cal');
# example 3
$p1 = popen('cal', 'r');
$s3 = stream_get_contents($p1);
# example 4
exec('cal', $a1);
# print
var_dump($s1, $s2, $s3, $a1);
