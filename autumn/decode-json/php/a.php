<?php
$s = '{"U2": {"Boy": ["Twilight", "I Will Follow"]}}';

# example 1
$o = json_decode($s);
$t = $o->U2->Boy[0];
var_dump($t == 'Twilight');

# example 2
$m = json_decode($s, true);
$t = $m['U2']['Boy'][0];
var_dump($t == 'Twilight');
