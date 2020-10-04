<?php
extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');
# example 1
$r_c = curl_init('https://speedtest.lax.hivelocity.net');
$r_f = fopen('index.html', 'w');
curl_setopt($r_c, CURLOPT_FILE, $r_f);
curl_exec($r_c);
fclose($r_f);
# example 2
$r_c = curl_init('https://speedtest.lax.hivelocity.net/10Mio.dat');
$r_f = fopen('10Mio.dat', 'w');
curl_setopt($r_c, CURLOPT_FILE, $r_f);
curl_setopt($r_c, CURLOPT_NOPROGRESS, false);
curl_exec($r_c);
fclose($r_f);
