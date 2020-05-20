<?php
extension_loaded('curl') or die('curl');
$s1 = 'http://speedtest.lax.hivelocity.net/100mb.file';
$r1 = curl_init($s1);
curl_setopt($r1, CURLOPT_HEADER, true);
curl_setopt($r1, CURLOPT_NOBODY, true);
curl_setopt($r1, CURLOPT_RETURNTRANSFER, true);
$s2 = curl_exec($r1);
$m1 = iconv_mime_decode_headers($s2);
print_r($m1);
