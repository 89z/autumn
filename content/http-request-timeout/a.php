<?php
$s1 = 'http://speedtest.lax.hivelocity.net:10';
$r1 = stream_context_create();
stream_context_set_option($r1, 'http', 'timeout', 5);
file_get_contents($s1, false, $r1);
