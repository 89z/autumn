<?php
$a1 = ['less', '-V'];
$a2 = array_map('escapeshellarg', $a1);
$s1 = implode(' ', $a2);
system($s1);
