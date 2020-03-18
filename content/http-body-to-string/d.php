<?php
$r1 = curl_init();
curl_setopt($r1, CURLOPT_RETURNTRANSFER, true);
curl_setopt($r1, CURLOPT_URL, 'http://speedtest.lax.hivelocity.net');
$s1 = curl_exec($r1);
curl_setopt($r1, CURLOPT_URL, 'http://speedtest.nyc.hivelocity.net');
$s2 = curl_exec($r1);
var_dump($s1, $s2);
