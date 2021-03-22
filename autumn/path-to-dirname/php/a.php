<?php
$s = 'C:\Windows\notepad.exe';

# example 1
$t = pathinfo($s, PATHINFO_DIRNAME);
var_dump($t == 'C:\Windows');

# example 2
$t = pathinfo($s)['dirname'];
var_dump($t == 'C:\Windows');
