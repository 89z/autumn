<?php
# example A
$sA = 'May
June';
# example B
$sB = "May
June";
# example C
$sC = <<<eof
May
June
eof;
# example D
$sD = <<<'eof'
May
June
eof;
# print
var_dump($sA, $sB, $sC, $sD);
