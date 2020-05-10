<?php
$s1 = '/sunday/monday.tar.xz';
# example 1
$s2 = dirname($s1);
# example 2
$s3 = pathinfo($s1, PATHINFO_DIRNAME);
# print
var_dump($s2 == '/sunday', $s3 == '/sunday');
