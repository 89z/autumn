<?php
$s1 = 'http://speedtest.lax.hivelocity.net/100mb.file';
$r1 = curl_init($s1);
curl_setopt($r1, CURLOPT_NOBODY, true);
curl_exec($r1);
$m1 = curl_getinfo($r1);
var_dump($m1);
