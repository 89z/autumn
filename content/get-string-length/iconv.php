<?php
function_exists('iconv_strlen') or die('php-iconv');
$n1 = iconv_strlen('♠');
var_dump($n1 == 1);
