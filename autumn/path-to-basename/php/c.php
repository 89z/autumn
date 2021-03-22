<?php
$f = new SplFileInfo('C:\Windows\notepad.exe');

# example 1
$s = $f->getBasename();
var_dump($s == 'notepad.exe');

# example 2
$s = $f->getBasename('.exe');
var_dump($s == 'notepad');
