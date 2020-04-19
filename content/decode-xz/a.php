<?php
# example 1
$s1 = shell_exec('xzdec a.tar.xz');
file_put_contents('a.tar', $s1);
# example 2
$s2 = shell_exec('xz -c -d a.tar.xz');
file_put_contents('b.tar', $s2);
# example 3
system('xz -d -k a.tar.xz');
