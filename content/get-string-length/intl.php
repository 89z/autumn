<?php
extension_loaded('intl') or die('intl');
$s = '📗';
$n = grapheme_strlen($s);
var_dump($n);
