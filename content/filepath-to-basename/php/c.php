<?php
$o = new SplFileInfo('C:\Windows\notepad.exe');
# example 1
$s1 = $o->getBasename();
# example 2
$s2 = $o->getBasename('.exe');
# print
var_dump($s1 == 'notepad.exe', $s2 == 'notepad');