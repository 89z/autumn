<?php
$s1 = 'https://api.github.com/rate_limit';
# example 1
$r1 = curl_init($s1);
curl_setopt($r1, CURLOPT_USERAGENT, PHP_VERSION);
curl_setopt($r1, CURLOPT_USERPWD, 'user:pass');
curl_exec($r1);
# example 2
$r2 = curl_init($s1);
curl_setopt($r2, CURLOPT_USERAGENT, PHP_VERSION);
curl_setopt($r2, CURLOPT_NETRC, true);
curl_exec($r2);
