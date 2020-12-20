<?php
extension_loaded('curl') or die('curl');

function http_head(string $s): bool {
   $r = curl_init($s);
   curl_setopt($r, CURLOPT_NOBODY, true);
   curl_exec($r);
   return curl_getinfo($r)['http_code'] == 200;
}

$b = http_head('http://speedtest.lax.hivelocity.net/10Mio.dat');
var_dump($b);
