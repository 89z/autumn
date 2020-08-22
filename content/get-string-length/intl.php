<?php
extension_loaded('intl') or die('intl');
$n = grapheme_strlen('📕');
var_dump($n);
