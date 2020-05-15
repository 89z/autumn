<?php
extension_loaded('intl') or die('extension=intl');
$s1 = '♠';
$n1 = grapheme_strlen($s1);
var_dump($n1);
