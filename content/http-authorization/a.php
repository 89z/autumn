<?php
$s1 = 'https://api.github.com/rate_limit';
# example 1
$s2 = getenv('USER');
$s3 = getenv('PASS');
$r1 = curl_init($s1);
curl_setopt($r1, CURLOPT_USERAGENT, 'anonymous');
curl_setopt($r1, CURLOPT_USERPWD, $s2 . ':' . $s3);
curl_exec($r1);
# example 2
$r2 = curl_init($s1);
curl_setopt($r2, CURLOPT_USERAGENT, 'anonymous');
curl_setopt($r2, CURLOPT_NETRC, true);
curl_exec($r2);
