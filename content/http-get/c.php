<?php
extension_loaded('curl') or die('curl');
$s = 'http://speedtest.lax.hivelocity.net';
# example 1
$r = curl_init($s);
curl_setopt($r, CURLOPT_RETURNTRANSFER, true);
echo curl_exec($r);
# example 2
$r_f = fopen('index.html', 'w');
$r_c = curl_init($s);
curl_setopt($r_c, CURLOPT_FILE, $r_f);
curl_setopt($r_c, CURLOPT_NOPROGRESS, false);
curl_exec($r_c);
fclose($r_f);
