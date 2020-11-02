<?php
$a = ['May', 'June'];
# example 1
$f = fn ($s) => print $s . "\n";
array_walk($a, $f);
# exmaple 2
$f = fn ($s, $n) => print $n . "\t" . $s . "\n";
array_walk($a, $f);
