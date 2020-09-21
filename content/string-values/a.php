<?php
# example 1
$s1 = 'One\One
One';
# example 2
$s2 = "Two
Two";
# example 3
$s3 = <<<eof
Three
Three
eof;
# example 4
$s4 = <<<'eof'
Four\Four
Four
eof;
# print
echo $s1, $s2, $s3, $s4, "\n";
