<?php
# example 1
system('ag');
# example 2
$a1 = ['ag', '-V'];
$a2 = array_map('escapeshellarg', $a1);
$s1 = implode(' ', $a2);
system($s1);
