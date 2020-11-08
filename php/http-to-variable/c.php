<?php
$s = 'https://speedtest.lax.hivelocity.net';
$r = fopen($s, 'r');
$s1 = stream_get_contents($r);
echo $s1;
