<?php
$s = '{"year": 2019, "month": 12}';
# example 1
$o1 = json_decode($s);
# example 2
$m2 = json_decode($s, true);
# example 3
$n3 = json_last_error();
# print
var_dump($o1, $m2, $n3 == JSON_ERROR_NONE);
