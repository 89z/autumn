<?php
$r = fopen('index.md', 'r');
$s = stream_get_contents($r);
var_dump($s);
