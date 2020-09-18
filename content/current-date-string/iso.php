<?php
# example 1
$s1 = date(DATE_W3C);
# example 2
$o = new DateTime;
$s2 = $o->format(DATE_W3C);
# example 3
$s3 = strftime('%FT%T');
# print
var_dump($s1, $s2, $s3);
