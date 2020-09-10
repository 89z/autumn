<?php
extension_loaded('curl') or die('curl');
extension_loaded('openssl') or die('openssl');
$r = curl_init('https://api.github.com/rate_limit');
curl_setopt($r, CURLOPT_USERAGENT, 'anonymous');
curl_setopt($r, CURLOPT_NETRC, true);
curl_setopt($r, CURLOPT_NETRC_FILE, getenv('USERPROFILE') . '/_netrc');
curl_exec($r);
