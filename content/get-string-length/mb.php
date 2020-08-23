<?php
extension_loaded('mbstring') or die('mbstring');
$s = '📕';
$n = mb_strlen($s);
var_dump($n);
