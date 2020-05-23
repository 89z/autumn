<?php
$s1 = 'http://speedtest.lax.hivelocity.net';
# example 1
$s2 = file_get_contents($s1);
# example 2
$r1 = fopen($s1, 'r');
$s3 = stream_get_contents($r1);
# print
var_dump($s2, $s3);
