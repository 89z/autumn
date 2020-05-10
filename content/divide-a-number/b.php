<?php
extension_loaded('gmp') or die('php-gmp');
$a1 = gmp_div_qr(19, 10);
var_dump($a1);
