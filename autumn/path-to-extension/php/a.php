<?php

# example 1
$s = pathinfo('file.php', PATHINFO_EXTENSION);
var_dump($s == 'php');

# example 2
$s = pathinfo('file.php')['extension'];
var_dump($s == 'php');
