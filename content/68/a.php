<?php
$s1 = 'C:\\Windows\\write.exe';
$s2 = pathinfo($s1, PATHINFO_DIRNAME);
var_dump($s2);
