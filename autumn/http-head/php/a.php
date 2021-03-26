<?php
extension_loaded('curl') or die('curl');

function http_head(string $s): bool {
   $c = curl_init($s);
   curl_setopt($c, CURLOPT_NOBODY, true);
   curl_exec($c);
   return curl_getinfo($c)['http_code'] == 200;
}

$b = http_head('http://speedtest.lax.hivelocity.net/10Mio.dat');
var_dump($b);
