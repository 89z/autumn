<?php
$s = 'http://speedtest.lax.hivelocity.net';
$f = fopen('index.html', 'w');
$c = curl_init($s);
curl_setopt($c, CURLOPT_FILE, $f);
curl_exec($c);
