<?php
extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');
$c = curl_init('https://api.github.com/search/repositories?q=database');
curl_setopt($c, CURLOPT_RETURNTRANSFER, true);
curl_setopt($c, CURLOPT_USERAGENT, 'anonymous');
$s = curl_exec($c);
echo $s;
