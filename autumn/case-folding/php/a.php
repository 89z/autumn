<?php
extension_loaded('mbstring') or die('mbstring');
# example 1
$s1 = mb_convert_case('March', MB_CASE_FOLD);
# example 2
$s2 = mb_convert_case('March', MB_CASE_FOLD_SIMPLE);
# print
var_dump($s1 == 'march', $s2 == 'march');
