<?php
date_default_timezone_set('America/Chicago');
$s = date(DATE_W3C, 1577858399);
var_dump($s == '2019-12-31T23:59:59-06:00');
