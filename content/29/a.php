<?php
$n = 365 * 24 * 60 * 60;
# example 1
$s = strftime('%a %b %#d %Y', $n);
# example 2
$s2 = strftime('%a %b %#e %Y', $n);
# print
var_dump($s == 'Fri Jan 1 1971', $s2 == 'Fri Jan 1 1971');
