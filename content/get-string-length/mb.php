<?php
extension_loaded('mbstring') or die('extension=mbstring');
$s1 = '♠';
$n1 = mb_strlen($s1);
var_dump($n1);
