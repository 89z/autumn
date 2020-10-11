<?php
extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');
$s = 'https://speedtest.lax.hivelocity.net';
$r = curl_init($s);
curl_setopt($r, CURLOPT_RETURNTRANSFER, true);
$s1 = curl_exec($r);
echo $s1;
