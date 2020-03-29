<?php
$s1 = 'http://speedtest.lax.hivelocity.net';
# example 1
$s2 = file_get_contents($s1);
# example 2
$r1 = fopen($s1, 'r');
$s3 = stream_get_contents($r1);
# example 3
$r2 = curl_init();
curl_setopt($r2, CURLOPT_RETURNTRANSFER, true);
curl_setopt($r2, CURLOPT_URL, $s1);
$s4 = curl_exec($r2);
# example 4
$r3 = curl_init($s1);
curl_setopt($r3, CURLOPT_RETURNTRANSFER, true);
$s5 = curl_exec($r3);
# print
var_dump($s2, $s3, $s4, $s5);
