<?php
$s1 = str_repeat('Sunday ', 2 ** 19);
# example 1
$t1 = microtime(true);
$a1 = explode(' ', $s1);
$t2 = microtime(true);
# example 2
$t3 = microtime(true);
$s2 = strtok($s1, ' ');
while ($s2 !== false) {
   $s2 = strtok(' ');
}
$t4 = microtime(true);
# print
echo $t2 - $t1, "\n"; # 0.046799898147583
echo $t4 - $t3, "\n"; # 0.015599966049194
