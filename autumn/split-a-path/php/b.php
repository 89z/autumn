<?php
$s = 'C:\Windows\notepad.exe';

# example 1
$m = pathinfo($s);
# array (
# 'dirname' => 'C:\\Windows',
# 'basename' => 'notepad.exe',
# 'filename' => 'notepad',
# 'extension' => 'exe',
# )
var_export($m);

# example 2
$t = pathinfo($s, PATHINFO_DIRNAME);
var_dump($t == 'C:\Windows');

# example 3
$t = pathinfo($s, PATHINFO_BASENAME);
var_dump($t == 'notepad.exe');

# example 4
$t = pathinfo($s, PATHINFO_EXTENSION);
var_dump($t == 'exe');
