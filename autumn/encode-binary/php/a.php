<?php
$s = "\n\v\f\r";
$t = base64_encode($s);
var_dump($t == 'CgsMDQ==');
