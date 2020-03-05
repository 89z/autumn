<?php
$s1 = 'http://speedtest.lax.hivelocity.net';
$s2 = file_get_contents($s1);
echo $s2;
