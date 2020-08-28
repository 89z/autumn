<?php
$s = '/aa/bb.cc.dd';
# example 1
$s1 = basename($s);
# example 2
$s2 = pathinfo($s, PATHINFO_BASENAME);
# example 3
$s3 = basename($s, '.dd');
# example 4
$s4 = pathinfo($s, PATHINFO_FILENAME);
# example 5
$s5 = pathinfo($s, PATHINFO_EXTENSION);
# print
var_dump(
   $s1 == 'bb.cc.dd',
   $s2 == 'bb.cc.dd',
   $s3 == 'bb.cc',
   $s4 == 'bb.cc',
   $s5 == 'dd'
);
