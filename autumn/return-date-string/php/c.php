<?php
date_default_timezone_set('America/Chicago');

# example 1
$s = date(DATE_W3C);
var_dump($s);

# example 2
$s = date('Y-m-d');
var_dump($s);
