<?php
extension_loaded('mbstring') or die('php-mbstring');
$s1 = 'SUNDAY';
# example 1
$s2 = mb_strtolower($s1);
# example 2
$s3 = mb_convert_case($s1, MB_CASE_LOWER);
# print
var_dump($s2, $s3);