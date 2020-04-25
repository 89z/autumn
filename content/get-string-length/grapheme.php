<?php
function_exists('grapheme_strlen') or die('php-intl');
$n1 = grapheme_strlen('♠');
var_dump($n1 == 1);
