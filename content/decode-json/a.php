<?php
$s = <<<eof
{"year": 2019, "month": 12, "day": 31}
eof;
# example 1
$o = json_decode($s);
# example 2
$m = json_decode($s, true);
# print
var_dump($o, $m);
