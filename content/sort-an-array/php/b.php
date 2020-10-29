<?php

$a = [
   ['year' => 2019, 'month' => 11],
   ['year' => 2018, 'month' => 12]
];

$f = fn ($m, $m2) => $m['year'] <=> $m2['year'];
usort($a, $f);
print_r($a);
