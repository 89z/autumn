<?php
$s = 'C:\Windows\notepad.exe';

# example 1
$t = pathinfo($s, PATHINFO_BASENAME);
var_dump($t == 'notepad.exe');

# example 2
$t = pathinfo($s, PATHINFO_FILENAME);
var_dump($t == 'notepad');
