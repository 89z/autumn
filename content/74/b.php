<?php
extension_loaded('gmp') or die('gmp');
$a = gmp_div_qr(29, 10);
var_dump($a);
