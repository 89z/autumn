<?php
extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');
$r = curl_init('https://speedtest.lax.hivelocity.net');
curl_setopt($r, CURLOPT_RETURNTRANSFER, true);
$s = curl_exec($r);
echo $s;