<?php
$r = fopen('http://speedtest.lax.hivelocity.net', 'r');
$s = stream_get_contents($r);
echo $s;
