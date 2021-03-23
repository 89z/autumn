<?php
$src = '{"U2": {"Boy": ["Twilight", "I Will Follow"]}}';

# example 1
$o = json_decode($src);
$dst = $o->U2->Boy[0];
var_dump($dst == 'Twilight');

# example 2
$m = json_decode($src, true);
$dst = $m['U2']['Boy'][0];
var_dump($dst == 'Twilight');
