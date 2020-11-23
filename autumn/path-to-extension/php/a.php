<?php
$s = 'a.php';
$s2 = pathinfo($s, PATHINFO_EXTENSION);
var_dump($s2 == 'php');
