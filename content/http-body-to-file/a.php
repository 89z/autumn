<?php
$s1 = 'http://speedtest.lax.hivelocity.net';
# example 1
copy($s1, 'a.html');
# example 2
$s2 = file_get_contents($s1);
file_put_contents('b.html', $s2);
