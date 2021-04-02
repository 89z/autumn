<?php
$s = 'https://api.github.com/rate_limit';
$c = curl_init($s);
curl_setopt($c, CURLOPT_USERAGENT, 'anonymous');
curl_exec($c);
