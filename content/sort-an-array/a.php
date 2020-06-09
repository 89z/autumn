<?php
$a = ['99', '100'];
# example 1
sort($a);
print_r($a);
# example 2
sort($a, SORT_NATURAL);
print_r($a);
# example 3
rsort($a);
print_r($a);
# example 4
rsort($a, SORT_NATURAL);
print_r($a);
