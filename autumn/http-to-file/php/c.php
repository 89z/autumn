<?php
extension_loaded('curl') or die('curl');
$in_r = curl_init('http://speedtest.lax.hivelocity.net/10Mio.dat');
$out_r = fopen('10Mio.dat', 'w');
curl_setopt($in_r, CURLOPT_FILE, $out_r);
curl_setopt($in_r, CURLOPT_NOPROGRESS, false);
curl_exec($in_r);
