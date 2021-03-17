<?php
$s = "\xa\xb\xc";
$t = base64_encode($s);
var_dump($t == 'CgsM');
