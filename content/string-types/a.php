<?php
# example 1
$s1 = "Sunday
Monday";
# example 2
$s2 = 'Sunday
Monday';
# example 3
$s3 = <<<eof
Sunday
Monday
eof;
# example 4
$s4 = <<<'eof'
Sunday
Monday
eof;
# print
var_dump($s1, $s2, $s3, $s4);
