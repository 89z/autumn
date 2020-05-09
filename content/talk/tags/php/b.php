<?php
$s1 = getenv('PUBLIC');
$s2 = realpath($s1);
var_dump($s2); # bool(false)
