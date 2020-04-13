<?php
$s1 = 'http://speedtest.lax.hivelocity.net/100mb.file';
# example 1
$m1 = get_headers($s1, 1);
# example 2
fopen($s1, 'r');
$a1 = $http_response_header;
# example 3
$r1 = fopen($s1, 'r');
$a2 = stream_get_meta_data($r1)['wrapper_data'];
# print
var_dump($m1, $a1, $a2);
