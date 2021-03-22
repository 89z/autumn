<?php
date_default_timezone_set('America/Chicago');

# example 1
$s = strftime('%F');
var_dump($s);

# example 2
$s = strftime('%FT%T');
var_dump($s);
