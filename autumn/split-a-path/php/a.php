<?php
$s = 'C:\Windows\notepad.exe';

# example 1
$t = basename($s);
var_dump($t == 'notepad.exe');

# example 2
$t = basename($s, '.exe');
var_dump($t == 'notepad');
