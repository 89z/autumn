<?php
$s = '{"U2": {"Boy": ["Twilight", "I Will Follow"]}}';

# example 1
$o = json_decode($s);
$s1 = $o->U2->Boy[0];

# example 2
$m = json_decode($s, true);
$s2 = $m['U2']['Boy'][0];

# print
var_dump($s1 == 'Twilight', $s2 == 'Twilight');
