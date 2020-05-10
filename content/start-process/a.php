<?php
# example 1
system('ag -V');
# example 2
$a1 = ['ag', '-V'];
$a2 = array_map('escapeshellarg', $a1);
$s1 = implode(' ', $a2);
system($s1);
# example 3
$a1 = [];
$a2 = [];
proc_open('sleep 10', $a1, $a2);
