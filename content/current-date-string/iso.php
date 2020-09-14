<?php
# example 1
$s = date(DATE_W3C);
# example 2
$o = new DateTime;
$s2 = $o->format(DATE_W3C);
# example 3
$s3 = strftime('%FT%T');
# print
var_dump($s, $s2, $s3);
