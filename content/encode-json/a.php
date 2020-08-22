<?php
$m['α/β'] = 10;
# example 1
$s1 = json_encode($m);
# example 2
$s2 = json_encode($m, JSON_PRETTY_PRINT);
# example 3
$s3 = json_encode($m, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
# print
var_dump($s1, $s2, $s3);
