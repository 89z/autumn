<?php
$a = ['/', '📗'];
# example 1
$s1 = json_encode($a);
# example 2
$s2 = json_encode($a, JSON_PRETTY_PRINT);
# example 3
$s3 = json_encode($a, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE);
# print
echo $s1, "\n", $s2, "\n", $s3, "\n";
