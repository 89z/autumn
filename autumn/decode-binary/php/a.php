<?php
$s = 'CgsMDQ==';
$t = base64_decode($s);
var_dump($t == "\xa\xb\xc\xd");
