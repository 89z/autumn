<?php
$s = 'https://speedtest.lax.hivelocity.net';
# example 1
$s1 = file_get_contents($s);
# example 2
$r = fopen($s, 'r');
$s2 = stream_get_contents($r);
# print
var_dump($s1, $s2);
