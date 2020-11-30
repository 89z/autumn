<?php

$a = [
   ['month' => 11, 'day' => 30],
   ['month' => 10, 'day' => 31]
];

$f = fn ($m, $m2) => $m['month'] <=> $m2['month'];
usort($a, $f);
print_r($a);
