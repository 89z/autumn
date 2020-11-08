<?php
$r = popen('php -v', 'r');
$s = stream_get_contents($r);
var_dump($s);
