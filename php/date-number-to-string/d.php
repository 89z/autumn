<?php
date_default_timezone_set('America/Chicago');
$n = 1577858399;
$s = strftime('%F', $n);
var_dump($s == '2019-12-31');
