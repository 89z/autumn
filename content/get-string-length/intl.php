<?php
extension_loaded('intl') or die('php-intl');
$n1 = grapheme_strlen('♠');
var_dump($n1 == 1);
