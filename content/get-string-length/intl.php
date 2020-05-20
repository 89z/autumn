<?php
extension_loaded('intl') or die('intl');
$s1 = '♠';
$n1 = grapheme_strlen($s1);
var_dump($n1);
