<?php
extension_loaded('mbstring') or die('php-mbstring');
$s1 = '♠♣♥♦';
# example 1
$s2 = mb_substr($s1, 0, 2);
# example 2
$s3 = mb_substr($s1, 2);
# example 3
$s4 = mb_substr($s1, -2);
# print
var_dump($s2 == '♠♣', $s3 == '♥♦', $s4 == '♥♦');
