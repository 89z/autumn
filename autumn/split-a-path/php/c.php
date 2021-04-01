<?php
$f = new SplFileInfo('C:\Windows\notepad.exe');
$s = $f->getBasename();
var_dump($s == 'notepad.exe');
