<?php
extension_loaded('curl') or die('curl');
$in_r = curl_init('http://speedtest.lax.hivelocity.net');
$out_r = fopen('index.html', 'w');
curl_setopt($in_r, CURLOPT_FILE, $out_r);
curl_exec($in_r);
