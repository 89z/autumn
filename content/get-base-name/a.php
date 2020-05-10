<?php
$s1 = '/sunday/monday.tar.xz';
# example 1
$s2 = basename($s1);
var_dump($s2 == 'monday.tar.xz');
# example 2
$s2 = pathinfo($s1, PATHINFO_BASENAME);
var_dump($s2 == 'monday.tar.xz');
# example 3
$s2 = pathinfo($s1, PATHINFO_FILENAME);
var_dump($s2 == 'monday.tar');
# example 4
$s2 = pathinfo($s1, PATHINFO_EXTENSION);
var_dump($s2 == 'xz');
# example 5
$m1 = pathinfo($s1);
$s2 = $m1['extension'];
var_dump($s2 == 'xz');
