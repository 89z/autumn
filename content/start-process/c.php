<?php
$a1 = [];
$a2 = [];
$r1 = proc_open('less index.md', $a1, $a2);
proc_close($r1);
