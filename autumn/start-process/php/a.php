<?php
$a = [];

# example 1
$p = proc_open('dust -V', $a, $a);
proc_close($p);

# example 2
$p = proc_open(['dust', '-V'], $a, $a);
proc_close($p);

# example 3
$p = proc_open('dust', $a, $a, '..');
proc_close($p);
