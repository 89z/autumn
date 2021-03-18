<?php
$s = "\xa\xb\xc\xd";
$t = base64_encode($s);
var_dump($t == 'CgsMDQ==');
