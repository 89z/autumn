<?php
function_exists('curl_init') or die('php-curl');
$s1 = 'http://speedtest.lax.hivelocity.net';
# example 1
$r1 = curl_init($s1);
curl_setopt($r1, CURLOPT_RETURNTRANSFER, true);
$s2 = curl_exec($r1);
# example 2
$r2 = curl_init();
curl_setopt($r2, CURLOPT_RETURNTRANSFER, true);
curl_setopt($r2, CURLOPT_URL, $s1);
$s3 = curl_exec($r2);
# print
var_dump($s2, $s3);
