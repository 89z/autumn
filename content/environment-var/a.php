<?php
$s1 = $_SERVER['BROWSER'];
$s2 = getenv()['BROWSER'];
$s3 = getenv('BROWSER');
# test
var_dump($s1, $s2, $s3);
