<?php
$a = [];

# example 1
proc_open('C:\Windows\notepad', $a, $a);

# example 2
$p = proc_open('C:\Windows\notepad', $a, $a);
proc_close($p);
