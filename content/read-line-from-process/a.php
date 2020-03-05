<?php
$p1 = popen('cal', 'r');
$s1 = fgets($p1);
var_dump($s1);
