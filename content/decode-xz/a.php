<?php
$s1 = 'python3-3.6.8-1.tar.xz';
# example 1
$s2 = shell_exec('xzdec ' . $s1);
file_put_contents('a.tar', $s2);
# example 2
$s3 = shell_exec('xz -c -d ' . $s1);
file_put_contents('b.tar', $s3);
# example 3
system('xz -d -k ' . $s1);
