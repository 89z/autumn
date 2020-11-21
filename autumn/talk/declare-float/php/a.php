<?php
declare(strict_types = 1);
error_reporting(E_ALL);
$n = 0x21 + 1_000 + 1e3;
var_dump($n == 2033);
