<?php
$s = 'http://speedtest.lax.hivelocity.net';
$c = curl_init($s);
curl_setopt($c, CURLOPT_NOPROGRESS, false);
curl_exec($c);
curl_close($c);
