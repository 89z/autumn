<?php
extension_loaded('curl') or die('curl');
$s1 = 'http://speedtest.lax.hivelocity.net/100mb.file';
# example 1
$r1 = curl_init($s1);
curl_setopt($r1, CURLOPT_HEADER, true);
curl_setopt($r1, CURLOPT_NOBODY, true);
curl_setopt($r1, CURLOPT_RETURNTRANSFER, true);
$s2 = curl_exec($r1);
# example 2
$r2 = curl_init($s1);
curl_setopt($r2, CURLOPT_NOBODY, true);
curl_exec($r2);
$m1 = curl_getinfo($r2);
# print
var_dump($s2, $m1);
