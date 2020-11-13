<?php
extension_loaded('curl') or die('curl');
$r = curl_init('http://speedtest.lax.hivelocity.net');
curl_setopt($r, CURLOPT_RETURNTRANSFER, true);
$s = curl_exec($r);
echo $s;
