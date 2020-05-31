<?php
extension_loaded('mbstring') or die('mbstring');
$s = '♠♣♥♦';
$s2 = mb_substr($s, 1, 1);
var_dump($s2 == '♣');
