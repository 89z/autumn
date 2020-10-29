<?php
$s = 'January';
$n = preg_match('/(?i)ja/', $s);
var_dump($n !== 0);
