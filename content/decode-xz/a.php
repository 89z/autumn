<?php
# example 1
system('xz -d -k a.tar.xz');
# example 2
$s1 = shell_exec('xz -c -d a.tar.xz');
file_put_contents('b.tar', $s1);
# example 3
$s2 = shell_exec('xzdec a.tar.xz');
file_put_contents('c.tar', $s2);
