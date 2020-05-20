<?php
extension_loaded('mbstring') or die('mbstring');
$s1 = '♠';
$n1 = mb_strlen($s1);
var_dump($n1);
