<?php
$s1 = getenv('PUBLIC');
$a1 = stat($s1);
var_dump($a1); # array(26)
