<?php
$r1 = fopen('a.txt', 'r');
$r2 = curl_init('smtps://smtp.gmail.com');
curl_setopt($r2, CURLOPT_MAIL_RCPT, ['monday@gmail.com']);
curl_setopt($r2, CURLOPT_NETRC, true);
curl_setopt($r2, CURLOPT_READDATA, $r1);
curl_setopt($r2, CURLOPT_UPLOAD, true);
curl_exec($r2);
