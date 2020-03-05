<?php
$s1 = 'http://speedtest.lax.hivelocity.net';
$r1 = fopen($s1, 'r');
$s2 = stream_get_contents($r1);
echo $s2;
