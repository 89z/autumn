<?php
declare(strict_types = 1);
error_reporting(E_ALL);
$m = getopt('a:b:', [], $n);
var_dump($m, $n);
