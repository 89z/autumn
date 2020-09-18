<?php
date_default_timezone_set('America/Chicago');
$n = 1577858399;
# example 1
$s1 = strftime('%a %b %#d %Y', $n);
# example 2
$s2 = strftime('%a %b %#e %Y', $n);
# print
var_dump($s1 == 'Tue Dec 31 2019', $s2 == 'Tue Dec 31 2019');
