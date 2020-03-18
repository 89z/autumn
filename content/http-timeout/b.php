<?php
$s1 = 'http://speedtest.lax.hivelocity.net:10';
$r1 = curl_init($s1);
curl_setopt($r1, CURLOPT_TIMEOUT, 5);
curl_exec($r1);
