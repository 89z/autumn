<?php
$s = 'http://speedtest.lax.hivelocity.net';
$c = curl_init($s);
curl_setopt($c, CURLOPT_RETURNTRANSFER, true);
echo curl_exec($c);
