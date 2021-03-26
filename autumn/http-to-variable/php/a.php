<?php
extension_loaded('curl') or die('curl');
$c = curl_init('http://speedtest.lax.hivelocity.net');
curl_setopt($c, CURLOPT_RETURNTRANSFER, true);
$s = curl_exec($c);
echo $s;
