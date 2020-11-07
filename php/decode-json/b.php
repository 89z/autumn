<?php
$s = '{"year": 2019, "month": 12}';
$o = json_decode($s);
$n = json_last_error();
var_dump($n == JSON_ERROR_NONE);
