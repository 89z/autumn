<?php
extension_loaded('mbstring') or die('php-mbstring');
$n1 = mb_strlen('♠');
var_dump($n1 == 1);
