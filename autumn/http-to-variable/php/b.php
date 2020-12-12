<?php
extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');
$r = curl_init('https://api.github.com/search/repositories?q=database');
curl_setopt($r, CURLOPT_RETURNTRANSFER, true);
curl_setopt($r, CURLOPT_USERAGENT, 'anonymous');
$s = curl_exec($r);
echo $s;
