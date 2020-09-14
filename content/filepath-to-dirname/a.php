<?php
$s = 'C:\\Windows\\write.exe';
$s2 = pathinfo($s, PATHINFO_DIRNAME);
var_dump($s2);
