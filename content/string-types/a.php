<?php
# example 1
$s1 = "sun mon tue
wed thu fri";
# example 2
$s2 = 'sun mon tue
wed thu fri';
# example 3
$s3 = <<<eof
sun mon tue
wed thu fri
eof;
# example 4
$s4 = <<<'eof'
sun mon tue
wed thu fri
eof;
# print
var_dump($s1, $s2, $s3, $s4);
