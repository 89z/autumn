<?php
extension_loaded('curl') or die('curl');
$curl = curl_init('http://speedtest.lax.hivelocity.net');
$open = fopen('index.html', 'w');
curl_setopt($curl, CURLOPT_FILE, $open);
curl_exec($curl);
