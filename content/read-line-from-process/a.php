<?php
$r1 = popen('ag -V', 'r');
$s1 = fgets($r1);
var_dump($s1);
