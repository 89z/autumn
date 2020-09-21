<?php
# example 1
$s1 = 'May
June';
# example 2
$s2 = "May
June";
# example 3
$s3 = <<<eof
May
June
eof;
# example 4
$s4 = <<<'eof'
May
June
eof;
# print
var_dump($s1, $s2, $s3, $s4);
