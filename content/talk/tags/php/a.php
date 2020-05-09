<?php
$s1 = getenv('PUBLIC');
$b1 = chdir($s1);
$s2 = getcwd();
var_dump($s2); # "/cygdrive/c/Users/Public"
