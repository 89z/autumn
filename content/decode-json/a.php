<?php
$s = <<<eof
{"Sunday": 10}
eof;
# example 1
$o = json_decode($s);
# example 2
$m = json_decode($s, true);
# print
var_dump($o, $m);
