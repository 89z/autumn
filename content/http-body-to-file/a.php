<?php
$s1 = 'http://speedtest.lax.hivelocity.net';
# example 1
copy($s1, 'sun.html');
# example 2
$s2 = file_get_contents($s1);
file_put_contents('mon.html', $s2);
