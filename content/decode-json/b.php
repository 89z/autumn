<?php
$s = '{"year": 2019}}';
json_decode($s);
$n = json_last_error();
var_dump($n != JSON_ERROR_NONE);
