<?php
$s = 'https://speedtest.lax.hivelocity.net';
# example 1
$s2 = file_get_contents($s);
# example 2
$r = fopen($s, 'r');
$s3 = stream_get_contents($r);
# print
var_dump($s2, $s3);
