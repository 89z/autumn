<?php
$s1 = 'http://speedtest.lax.hivelocity.net/100mb.file';
$r1 = fopen('100mb.file', 'w');
$r2 = curl_init();
curl_exec($r2);
curl_setopt($r2, CURLOPT_FILE, $r1);
curl_setopt($r2, CURLOPT_NOPROGRESS, false);
