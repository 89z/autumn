<?php
extension_loaded('curl') or die('curl');

$curl = curl_init('http://speedtest.lax.hivelocity.net/10Mio.dat');
$open = fopen('10Mio.dat', 'w');
curl_setopt($curl, CURLOPT_FILE, $open);
curl_setopt($curl, CURLOPT_NOPROGRESS, false);
curl_exec($curl);
