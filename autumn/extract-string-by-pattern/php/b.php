<?php
$s = 'south north';
preg_match('/o(..)/', $s, $a);
var_dump($a[1] == 'ut');
