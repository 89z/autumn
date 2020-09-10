<?php
$s = '{"year": 2020}}';
json_decode($s);
$n = json_last_error();
var_dump($n != JSON_ERROR_NONE);
