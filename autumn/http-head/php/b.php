<?php

function http_head(string $s): bool {
   return get_headers($s)[0] == 'HTTP/1.1 200 OK';
}

$b = http_head('http://speedtest.lax.hivelocity.net/10Mio.dat');
var_dump($b);
