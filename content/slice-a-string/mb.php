<?php
extension_loaded('mbstring') or die('php-mbstring');
$s1 = '♠♣♥♦';
$s2 = mb_substr($s1, 1, 2);
var_dump($s2);
