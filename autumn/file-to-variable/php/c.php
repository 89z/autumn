<?php
$f = fopen('a.php', 'r');
$s = stream_get_contents($f);
var_dump($s);
