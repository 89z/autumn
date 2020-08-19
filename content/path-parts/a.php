<?php
$s1 = '/a/b.tar.xz';
# example 1
$s2 = basename($s1);
var_dump($s2 == 'b.tar.xz');
# example 2
$s2 = basename($s1, '.xz');
var_dump($s2 == 'b.tar');
# example 3
$s2 = pathinfo($s1, PATHINFO_BASENAME);
var_dump($s2 == 'b.tar.xz');
# example 4
$s2 = pathinfo($s1, PATHINFO_FILENAME);
var_dump($s2 == 'b.tar');
# example 5
$s2 = pathinfo($s1, PATHINFO_EXTENSION);
var_dump($s2 == 'xz');
