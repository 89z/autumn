<?php
$a1 = ['Sun', 'Mon'];
# example 1
$s1 = json_encode($a1);
# example 2
$s2 = json_encode($a1, JSON_UNESCAPED_SLASHES);
# example 3
$s3 = json_encode($a1, JSON_PRETTY_PRINT | JSON_UNESCAPED_UNICODE);
# print
var_dump($s1, $s2, $s3);
