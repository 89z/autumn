<?php
$s1 = 'Sunday';
preg_match('/(.*)day/', $s1, $a1);
$s2 = $a1[1];
var_dump($s2);
