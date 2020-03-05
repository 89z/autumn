<?php
$s1 = 'http://speedtest.lax.hivelocity.net';
$r1 = curl_init($s1);
curl_setopt($r1, CURLOPT_RETURNTRANSFER, true);
$s2 = curl_exec($r1);
echo $s2;
