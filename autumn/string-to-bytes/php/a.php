<?php
$s = 'CgsMDQ==';
$t = base64_decode($s);
var_dump($t == "\n\v\f\r");
